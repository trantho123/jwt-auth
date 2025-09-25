FROM golang:alpine3.22 AS builder
WORKDIR /app
COPY ./api .
RUN apk update

RUN go build -o main ./cmd/main.go

FROM alpine:3.22.1
WORKDIR /app
COPY --from=builder /app/main .
COPY ./api/.env .
CMD [ "./main" ]
