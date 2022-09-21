FROM golang:1.17-alpine

WORKDIR /usr/app

COPY . /usr/app/

RUN apk add --update make

RUN make vendor

RUN make build

EXPOSE 8080

CMD ["make","run"]
