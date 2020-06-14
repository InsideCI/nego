# first stage: golang base image as a builder
FROM golang:alpine AS builder

ENV CGO_ENABLED=0

LABEL maintainer "Cleanderson Lins <cleandersonlins@gmail.com>"

# install git and ca-certificates
RUN apk add --no-cache ca-certificates git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o nego .

FROM alpine:latest AS final

WORKDIR /root/

COPY --from=builder /app/nego .
COPY --from=builder /app/.env .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs

EXPOSE 80
EXPOSE 8081

VOLUME ["/cert-cache"] 

# RUN chmod +x .env

# ENTRYPOINT ["/app"]
CMD ["./nego", "-port", "8081" ,"-prod"]

