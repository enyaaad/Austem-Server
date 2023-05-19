FROM golang:1.20.4 as builder

RUN go version


COPY ./ /go/src/github.com/enyaaad/Austem-Server/
WORKDIR /go/src/github.com/enyaaad/Austem-Server/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

RUN apk update
RUN apk add git
RUN apk upgrade
RUN apk --no-cache add curl

COPY --from=0  /go/src/github.com/enyaaad/Austem-Server/.bin/app .
COPY --from=0  /go/src/github.com/enyaaad/Austem-Server/config/ ./config/

CMD ["./app"]