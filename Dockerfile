FROM golang:1.15.6-alpine

COPY . .

ENV GOPATH=

ENV CGO_ENABLED=0

RUN go build main.go

ENV GIN_MODE=release

EXPOSE 80

CMD ["/go/main"]
