FROM golang

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN go get -d -v ./...

RUN go install -v ./...