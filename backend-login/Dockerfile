FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
COPY main.go ./
COPY handlers ./handlers
COPY utils ./utils

RUN go build -o backend .

EXPOSE 8080

CMD ["./backend"]
