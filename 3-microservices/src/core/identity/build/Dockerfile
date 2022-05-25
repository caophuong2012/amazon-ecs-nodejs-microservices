FROM golang:1.17.6-alpine3.15

WORKDIR /app

COPY . .

ENV GO111MODULE=on

RUN go mod vendor
RUN go mod verify

RUN go build -o /api cmd/handlers/main.go

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

CMD /wait && /api
EXPOSE 3000