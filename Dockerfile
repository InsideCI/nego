FROM golang:latest

LABEL maintainer "Cleanderson Lins <cleandersonlins@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build .

EXPOSE 8081

CMD ["./nego -PORT 8080"]

