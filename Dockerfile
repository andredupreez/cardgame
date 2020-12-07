FROM golang:1.15-alpine AS builder

WORKDIR /go/src/github.com/andredupreez/cardgame
COPY . /go/src/github.com/andredupreez/cardgame
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/poker

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/poker .
CMD ["./poker"]  