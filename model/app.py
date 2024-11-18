from fastapi import FastAPI, File, UploadFile, HTTPException
from fastapi.responses import JSONResponse
import os
import pathlib
import json
import shutil
from math import ceil
import torch
import torchaudio
from time import gmtime, strftime
from src import constants, model

app = FastAPI()

# Настройки и инициализация
CWD = os.getcwd()
pathlib.Path(CWD).joinpath('data/processed').mkdir(parents=True, exist_ok=True)
pathlib.Path(CWD).joinpath('saved_weights').mkdir(parents=True, exist_ok=True)

MODEL_SAMPLE_RATE = constants.MODEL_SAMPLE_RATE
CHUNK_SIZE = constants.CHUNK_SIZE
BATCH_SIZE = constants.BATCH_SIZE
CHUNKS_WITH_WOLF_RATE = constants.CHUNKS_WITH_WOLF_RATE
DEVICE = torch.device("cuda:0" if torch.cuda.is_available() else "cpu")
model_instance = model.WolfClassifier().to(DEVICE)

@app.post("/process-audio/")
async def process_audio(file: UploadFile = File(...), confidence_threshold: float = 0.5):
    """
    Обрабатывает аудиофайл для обнаружения воя волков.

    :param file: Загружаемый WAV-файл.
    :param confidence_threshold: Минимальная уверенность модели для разметки (0.0 - 1.0).
    :return: Архив с размеченными данными.
    """
    if not file.filename.endswith(".wav"):
        raise HTTPException(status_code=400, detail="Только WAV-файлы поддерживаются.")

    # Сохранение загруженного файла
    input_path = pathlib.Path(CWD).joinpath(f"data/processed/{file.filename}")
    with open(input_path, "wb") as buffer:
        shutil.copyfileobj(file.file, buffer)

    # Обработка файла
    waveform, sample_rate = torchaudio.load(input_path)
    if sample_rate != MODEL_SAMPLE_RATE:
        waveform = torchaudio.functional.resample(waveform, sample_rate, MODEL_SAMPLE_RATE)

    waveform = waveform.mean(dim=0, keepdim=True)
    audio_chunks = list(torch.split(waveform, MODEL_SAMPLE_RATE * CHUNK_SIZE, dim=1))

    current_file_markup = []
    duration = []
    for start_index in range(0, len(audio_chunks), BATCH_SIZE):
        current_tensors = audio_chunks[start_index : start_index + BATCH_SIZE]
        duration.extend([CHUNK_SIZE] * (len(current_tensors) - 1))
        duration.append(ceil(current_tensors[-1].shape[1] / MODEL_SAMPLE_RATE))

        # Дополнение последнего чанка до нужного размера
        if current_tensors[-1].shape[1] < CHUNK_SIZE * MODEL_SAMPLE_RATE:
            padding = torch.nn.ConstantPad1d(
                padding=(0, CHUNK_SIZE * MODEL_SAMPLE_RATE - current_tensors[-1].shape[1]), value=0
            )
            current_tensors[-1] = padding(current_tensors[-1])

        batch = torch.stack(current_tensors, dim=0).squeeze(1).to(DEVICE)
        current_file_markup.extend(list(model_instance.get_wolf_probability(batch).cpu().numpy()))

    # Формирование разметки
    markup = {}
    running_duration = 0.0
    number_of_chunks_with_noise = 0
    for duration_value, wolf_probability in zip(duration, current_file_markup):
        start = running_duration
        end = running_duration + duration_value

        if wolf_probability > confidence_threshold:
            number_of_chunks_with_noise += 1
            if file.filename not in markup:
                markup[file.filename] = []
            markup[file.filename].append(
                (strftime("%H:%M:%S", gmtime(start)), strftime("%H:%M:%S", gmtime(end)))
            )

        running_duration += duration_value

    return JSONResponse(content=markup)
