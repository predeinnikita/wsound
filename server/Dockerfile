FROM golang:1.22 AS builder

WORKDIR /server
COPY . .

RUN go mod download
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o main cmd/server/main.go

FROM alpine:3.15
WORKDIR /root/

COPY --from=builder server/main .

RUN chmod +x main

EXPOSE 8080
CMD ["./main"]

