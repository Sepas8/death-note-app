FROM golang:1.24.3 AS base

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o backend-avanzada.o

EXPOSE 8000

CMD ["/build/backend-avanzada.o"]