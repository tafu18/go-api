FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# cmd/api altındaki main.go'yu derle
RUN go build -o main ./cmd/api/main.go

EXPOSE 8080

CMD ["./main"]
