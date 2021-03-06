FROM golang:latest

COPY . /src

ADD . /go/src
WORKDIR /go/src


RUN apt-get update
RUN apt-get install vim -y
RUN go get -u github.com/go-sql-driver/mysql
RUN go build github.com/go-sql-driver/mysql

# COPY . ./

# RUN apt-get update
# RUN apt-get install vim -y

# RUN mkdir -p /go/src/github.com/pavel/gorm-psql
# WORKDIR /go/src/github.com/pavel/gorm-psql

# ADD . /go/src/github.com/pavel/gorm-psql

# RUN go get -v
# RUN go get -u github.com/gin-gonic/gin
# RUN go get -u github.com/go-gorm/gorm