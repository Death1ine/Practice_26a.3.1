# Этап 1: Сборка приложения
FROM golang:1.23.1 AS build

# Включение поддержки модулей
ENV GO111MODULE=on

# Установка рабочей директории
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Установка зависимостей
RUN go mod download

# Копируем весь исходный код
COPY . .

# Сборка приложения
RUN go build -o pipeline-app .

# Этап 2: Минимальный образ для запуска
FROM alpine:3.18

# Установка сертификатов для HTTPS
RUN apk add --no-cache ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем собранный бинарник из первого этапа
COPY --from=build /app/pipeline-app .

# Открываем порт для приложения
EXPOSE 8080

# Указываем команду для запуска приложения
CMD ["./pipeline-app"]
