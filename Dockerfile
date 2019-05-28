FROM golang:1.11.10-alpine3.9 AS build
RUN apk --no-cache add gcc g++ make ca-certificates

WORKDIR /go/src

COPY vendor/ ./

WORKDIR /go/src/Phonebook

COPY util util
COPY data data
COPY db db
COPY handlers handlers
COPY schema schema
COPY logger logger
COPY Book.go Book.go


RUN go install ./...

FROM alpine:3.9
RUN apk --no-cache add curl
EXPOSE 8080/tcp
WORKDIR /usr/bin
COPY --from=build /go/bin .
