FROM golang:1.20.7-alpine
WORKDIR /app

COPY . .

RUN go mod download

CMD go run main.go migrate up; go run main.go server
