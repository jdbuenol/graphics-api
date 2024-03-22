FROM golang:1.22

WORKDIR /app 

COPY . .

RUN go mod download

EXPOSE 8080

RUN go build -o main ./cmd/server/main.go

CMD ["./main"]
