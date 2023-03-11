FROM golang:1.20.2-alpine

WORKDIR /usr/app

COPY . /usr/app/

RUN export GOPROXY=direct

RUN go build -o nymeria ./cmd/nymeria/main.go

EXPOSE 9898

CMD ["./nymeria"]