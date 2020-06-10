# first stage: golang base image as a builder
FROM golang:latest-alpine AS builder

LABEL maintainer "Cleanderson Lins <cleandersonlins@gmail.com>"

# install git and ca-certificates
RUN apk add --no-cache ca-certificates git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build .

FROM alpine:latest AS final

RUN apk add --no-cache ca-certificates

COPY --from=builder /app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs

EXPOSE 80
EXPOSE 8081

VOLUME ["/cert-cache"]

CMD ["./nego -port 8081 -prod"]

