FROM golang:1.21.1-alpine as builder

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0

RUN go build -o main

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

ENV GIN_MODE=release

EXPOSE 80

CMD ["./main"]
