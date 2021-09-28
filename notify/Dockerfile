# syntax=docker/dockerfile:1

##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app ./

USER nonroot:nonroot

ENTRYPOINT ["./main"]
