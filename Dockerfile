FROM golang:1.21.1-alpine as builder

COPY . .

ENV CGO_ENABLED=0

ENV GOPATH=

RUN go build -o main

RUN go test

FROM scratch

COPY --from=builder /go/main .

ENV GIN_MODE=release

EXPOSE 80

CMD ["./main"]
