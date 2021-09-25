FROM golang:1.17-alpine AS build

WORKDIR /go/src/github.com/BenShaffer/api

COPY go.mod ./
COPY go.sum ./

RUN go mod download
ADD . .
RUN go build -o api ./cmd/api.go

FROM alpine:latest AS app

WORKDIR /go/src/github.com/BenShaffer/api
COPY --from=build /go/src/github.com/BenShaffer/api .

ENTRYPOINT [ "./api" ]