FROM golang:1.23-alpine 

WORKDIR /app

COPY . /app/

RUN export GOPROXY=direct

EXPOSE 9898

# install make, psql
RUN apk add --no-cache make postgresql-client
RUN make build

CMD ["./nymeria"]