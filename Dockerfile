########### builder ###########
# AS builder
FROM golang:1.15.6 

RUN go get github.com/gin-gonic/gin

COPY . .

ENV CGO_ENABLED=0

RUN go build main.go

ENV GIN_MODE=release

EXPOSE 80

CMD ["/go/main"]

########### gin-key-value-store ###########
# FROM scratch

# COPY templates .

# COPY --from=builder /go/main .

# ENV GIN_MODE=release

# EXPOSE 80

# CMD ["/main"]
