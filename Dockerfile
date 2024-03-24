FROM golang:alpine3.19

RUN apk update
COPY ./api /app
WORKDIR /app

RUN go build -o main cmd/main.go

CMD ["./main"]