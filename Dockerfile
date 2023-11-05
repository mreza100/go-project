FROM golang:1.20-alpine AS builder

ENV APP_HOME /usr/app

WORKDIR "$APP_HOME"
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o waringin

FROM golang:1.20-alpine AS runner

ENV APP_HOME /usr/app
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

RUN addgroup --system --gid 1001 go
RUN adduser --system --uid 1001 gouser

COPY --from=builder --chown=gouser:go "$APP_HOME"/.env $APP_HOME
COPY --from=builder --chown=gouser:go "$APP_HOME"/waringin $APP_HOME

EXPOSE 6000

CMD ["./waringin"]
