FROM golang:1.15.6-alpine as builder

COPY go.mod  go.sum  main.go ./

RUN CGO_ENABLED=0 GOPATH= go build main.go

FROM alpine

WORKDIR /opt/app

COPY --from=builder /go/main /opt/app

COPY ./templates /opt/app/templates

ENV GIN_MODE=release

EXPOSE 80

CMD ["./main"]
