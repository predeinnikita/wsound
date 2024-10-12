const express = require("express");
const router = express.Router();
const multer = require("multer");
const mm = require("music-metadata");

// Конфигурация multer для загрузки файла
const storage = multer.memoryStorage(); // Файл будет храниться в памяти
const upload = multer({ storage: storage });

router.post(
  "/info",
  upload.single("audioFile"),
  async function (req, res, next) {
    try {
      if (!req.file) {
        return res.status(400).send("Файл не загружен");
      }

      const loadMusicMetadata = await mm.loadMusicMetadata();

      // Получение метаданных из загруженного файла
      const metadata = await loadMusicMetadata.parseBuffer(req.file.buffer);
      const { format, common } = metadata;
      const audioInfo = {
        duration: format.duration, // Длительность в секундах
        bitrate: format.bitrate, // Битрейт
        sampleRate: format.sampleRate, // Частота дискретизации
        codec: format.codec, // Кодек
        title: common.title || "Не указано", // Название трека (если есть метаданные)
        artist: common.artist || "Не указано", // Исполнитель (если есть метаданные)
      };

      res.json(audioInfo);
    } catch (error) {
      console.error(error);
      res.status(500).send("Ошибка при обработке аудиофайла");
    }
  }
);

module.exports = router;
module.exports = router;
