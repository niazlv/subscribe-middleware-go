# стадия 1. Сборка проекта(не будет включен в финальный образ)
FROM golang:1.21-alpine as builder

# Ставим какие-либо зависимости системы
# Устанавливаем необходимые зависимости для CGO
# RUN apk add --no-cache gcc musl-dev
RUN apk add --no-cache gcc=13.2.1_git20231014-r0 musl-dev=1.2.4_git20230717-r4

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /build

# Копируем модули Go(отдельно от исходников, чтобы лучше кешировалось и каждый раз по новой не обновлялось)
COPY go.mod .
COPY go.sum .

# Скачиваем зависимости уже GO
RUN go mod download

# Копируем исходный код приложения
COPY . .

# Собираем приложение с поддержкой CGO
ENV CGO_ENABLED=1
RUN go build -o main ./cmd/subscribe-middleware-go/main.go 


# стадия 2. Запуск в отдельном контейнере
FROM alpine:3.14

# Создаем директорию, в которой в последующем будем работать и хранить бинарник
RUN mkdir /app

# костыль доступа к storage host машины
ARG UID=1000
ARG GID=1000

# создаем пользователя, под которым будет выполняться код и его папку, в которой он сможет писать
RUN adduser -D --gid ${GID} --uid ${UID} appuser && mkdir /app-data && chown -R appuser /app-data

# создаем папку для работы программы(chown можно оптимизировать с прошлой командой)
RUN mkdir /app-data/storage && chown -R appuser /app-data/storage

# линкуем сторж, к папке, в которой пользователь может писать(потому что лень переписывать код программы, указывая другой патч)
RUN ln -s /app-data/storage /app/storage

WORKDIR /app

# Копируем файл из под прошлого этапа, для пользователя root("защита" от перезаписи запускаемого бинарника)
COPY --from=builder --chown=root:root /build/main .

# На всякий выдаем права на чтение и выполнение другим(не root)
RUN chmod 755 main

# перевести в режим продакшена
# ENV GIN_MODE=release

# понижаем права, для запуска кода(чтобы код был запущен без ЛИШНИХ привелегий)
USER appuser

# Указываем команду для запуска приложения
CMD ["/app/main"]