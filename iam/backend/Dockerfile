FROM golang:alpine3.21

ENV SERVERPORT=3000
WORKDIR /app

COPY *.go .
COPY go.mod .
COPY go.sum .

RUN go build

CMD [ ".\iam" ]
