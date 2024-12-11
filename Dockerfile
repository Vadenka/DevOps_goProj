# Используем минимальный образ с Go
FROM golang:1.20-alpine

# Установка зависимостей
RUN apk add --no-cache git

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем все файлы в контейнер
COPY . .

# Загружаем зависимости
RUN go mod init example.com/m/v2

# Сборка приложения
RUN go build -o app .

# Команда для запуска
CMD ["./app"]
