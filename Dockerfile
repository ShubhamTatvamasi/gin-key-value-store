########### builder ###########
FROM golang:1.15.6 AS builder

RUN go get github.com/gin-gonic/gin

COPY main.go .

ENV CGO_ENABLED=0

RUN go build main.go

########### gin-key-value-store ###########
FROM scratch

COPY --from=builder /go/main .

ENV GIN_MODE=release

EXPOSE 80

CMD ["/main"]
