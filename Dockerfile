FROM golang:1.23-alpine AS base

WORKDIR /app

ENV GOPROXY=direct

RUN apk add --no-cache make postgresql-client git curl

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

# Development Mode
FROM base AS dev
COPY . .
RUN make install-tools
EXPOSE 9898
CMD ["make", "dev"]

# Production Mode
FROM base AS prod
COPY . .
RUN make build
EXPOSE 9898
CMD ["./build/nymeria"]