# стадия 1. Сборка проекта(не будет включен в финальный образ)
FROM golang:1.21-alpine as builder

# Ставим какие-либо зависимости системы
# Устанавливаем необходимые зависимости для CGO
RUN apk add --no-cache gcc musl-dev

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /build

# Копируем модули Go(отдельно от исходников, чтобы лучше кешировалось и каждый раз по новой не обновлялось)
ADD go.mod .
ADD go.sum .

# Скачиваем зависимости уже GO
RUN go mod download

# Копируем исходный код приложения
COPY . .

# Собираем приложение с поддержкой CGO
ENV CGO_ENABLED=1
RUN go build -o main ./cmd/subscrbe-middleware-go/main.go 


# стадия 2. Запуск в отдельном контейнере
FROM alpine:3.14

WORKDIR /app

RUN mkdir storage
COPY --from=builder /build/main .

# перевести в режим продакшена
# ENV GIN_MODE=release

# Указываем команду для запуска приложения
CMD ["/app/main"]