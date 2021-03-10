FROM golang:alpine as builder

WORKDIR /data
COPY . .

RUN go build -o worker .

FROM alpine as app

COPY --from=builder /data/worker /data/worker
WORKDIR /data/logs

CMD [ "/data/worker" ]