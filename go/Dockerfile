FROM golang:latest


WORKDIR /go/src/app
COPY . .

RUN go env -w GO111MODULE=auto
RUN go get -d -v ./...
RUN go install -v ./...


RUN apt-get update
RUN apt-get install vim -y
RUN go get -u github.com/go-gorm/gorm
