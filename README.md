# Сервис распознавания воя волков на аудиозаписях

## Локальный запуск

### Frontend
Необходимо: NodeJS (LTS)
```bash
    cd client-v2 && npm ci && npm start
```

### Backend 
Необходимо: Docker Compose
```bash
    docker-compose up
```

## Подготовка к деплою

### Frontend
Для сборки клиента нужно выполнить команду:
```bash
    cd client-v2 && npm ci && npm run build
```
Сборка будет лежать в директории client-v2/build

### Модель распознавания

Сборка образа модели:
```bash
    cd model && docker build . -t wsound_model
```

Запуск образа модели:
```bash
    docker run --env-file .env wsound_model
```

### Backend

Сборка образа Backend:
```bash
    cd server && docker build . -t wsound_server
```

Запуск образа:
```bash
docker run --env-file .env wsound_server
```

Также для работы сервера необходимо развернуть базу данных Postgres и Minio (https://min.io/docs/minio/kubernetes/upstream/index.html).
Реквизиты для подключения к Postgres и Minio настраиваются в файле /server/.env
