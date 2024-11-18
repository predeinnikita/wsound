FROM pytorch/pytorch:2.2.1-cuda11.8-cudnn8-runtime

WORKDIR /app

RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get install -y ffmpeg

COPY ./requirements.txt ./install/requirements.txt
RUN pip install -r ./install/requirements.txt

EXPOSE 8000

COPY ./app.py ./app.py
COPY ./src ./src

CMD ["fastapi", "run", "app.py"]
