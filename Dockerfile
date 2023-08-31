FROM golang:1.21.0-alpine as builder

COPY . .

ENV CGO_ENABLED=0

ENV GOPATH=

RUN go build -o main

RUN go test

FROM alpine

WORKDIR /opt/app

COPY --from=builder /go/main /opt/app

COPY ./templates /opt/app/templates

ENV GIN_MODE=release

EXPOSE 80

CMD ["./main"]
