FROM golang:1.14


WORKDIR /go/src/app
COPY . .

RUN go env -w GO111MODULE=auto
RUN apt-get update
RUN apt-get install vim -y
WORKDIR /go
RUN export GOROOT=/usr/lib/go
RUN export GOPATH=$HOME/go
RUN export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/mysql
RUN go get -d -v ./...
RUN go install -v ./...

WORKDIR /go/src
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/mysql