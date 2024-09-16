# Avito Test Project Fall 2024

Проект разработан на языке Go. Для быстрого старта нужно иметь установленный Docker и docker-compose и следовать инструкциям, описанным ниже

## Быстрый старт:

1. Склонируйте репозиторий:
```bash
git clone https://github.com/tlb-katia/Avito_RestAPI
```

2. Перейдите в директорию проекта и запустите докер образ:
```bash
make docker-up
```

3. После сборки проект будет принимать запросы по адресу:
```bash
http://localhost:8080/api/ping
```

4. Также проект запускается по адресу:
```bash
https://cnrprod1725723419-team-78602-32501.avito2024.codenrock.com/api/ping
```


## Старт проекта, который залит в деплой:

1. Перейдите в папку:
```bash
cd deploy/zadanie-6105/
```

2. Соберите проект с помощью Docker:
```bash
docker build . -t avito-tender-service
```

3. Запустите проект:
```bash
docker run -p 8080:8080 avito-tender-service
```
