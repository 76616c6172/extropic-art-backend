# Dockerfile for project exia (includes webserver and assets)
# for local testing only

# Stage 1, build the binary in a container
###
FROM golang:1.18.3-alpine3.16 as builder
RUN apk add build-base
WORKDIR /src

## Get dependencies
COPY ./go.mod ./
RUN go mod download
COPY . .

## Compile static binary
WORKDIR /src/CONTROLLER
RUN go build -ldflags '-linkmode external -extldflags "-fno-PIC -static"' -o CONTROLLER

# Stage 2, build the final container with the minimum required files
###
FROM alpine as release
WORKDIR /src
#COPY . .
COPY --from=builder /src/CONTROLLER/CONTROLLER /src/CONTROLLER/CONTROLLER
COPY --from=builder /src/model /src/
COPY --from=builder /src/view /src/

EXPOSE 8080
WORKDIR /src/controller
RUN mkdir logs
ENTRYPOINT ["/src/CONTROLLER/CONTROLLER"]