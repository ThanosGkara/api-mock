FROM golang:1.17.2-alpine3.14 as builder

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache git

# setting working directory
WORKDIR /go/src/app

# installing dependencies
RUN go mod init
RUN go get github.com/sirupsen/logrus
RUN go get github.com/buaazp/fasthttprouter
RUN go get github.com/valyala/fasthttp

COPY / /go/src/app/
RUN go build -o mock

FROM alpine:3.14

WORKDIR /go/src/app
COPY --from=builder /go/src/app/mock /go/src/app/mock

COPY *.json /app/mocks/

EXPOSE 80

CMD ["./mock"]