FROM golang:1.18 AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0 
RUN go build -o main .
ENTRYPOINT ["./main"]