FROM golang:1.24-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Открываем порт
EXPOSE 8080

# Запускаем приложение напрямую (код будет монтироваться через volume)
# Используем -mod=mod чтобы игнорировать vendor директорию
CMD ["go", "run", "-mod=mod", "cmd/main.go"]
