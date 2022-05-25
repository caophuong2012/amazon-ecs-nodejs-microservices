FROM golang:1.17.6-alpine3.15

WORKDIR /app

COPY . .

ENV GO111MODULE=on

RUN go mod vendor
RUN go mod verify

RUN go build -o /api cmd/handlers/main.go

CMD /api