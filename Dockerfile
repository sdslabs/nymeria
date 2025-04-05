FROM golang:1.23-alpine 

WORKDIR /usr/app

COPY . /usr/app/

RUN export GOPROXY=direct

RUN go build -o nymeria ./cmd/nymeria/main.go

EXPOSE 9898

CMD ["./nymeria"]