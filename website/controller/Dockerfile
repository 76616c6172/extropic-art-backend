# Dockerfile for exia-controller (Go webserver)
# Stage 1, build the binary in a container
FROM golang:1.17.7-alpine as builder
RUN apk add build-base
WORKDIR /src

## Get dependencies
COPY go.mod ./
RUN go mod download
RUN touch test
COPY . .

## Compile static binary
RUN go build -ldflags '-linkmode external -extldflags "-fno-PIC -static"' -o exia-controller

# Stage 2, build the final container with the minimum required files
FROM alpine as release
WORKDIR /src
COPY . .
COPY --from=builder /src/exia-controller /src/
COPY --from=builder /src/model /src/
COPY --from=builder /src/view /src/

EXPOSE 8080
ENTRYPOINT ["/src/exia-controller"]