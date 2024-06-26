FROM golang:1.22.4-alpine as builder

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0

RUN go build -o main

FROM scratch

COPY --from=builder /app/main .

ENV GIN_MODE=release

EXPOSE 80

CMD ["./main"]
