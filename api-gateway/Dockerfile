FROM golang:1.18.2-alpine3.15 AS builder

RUN mkdir app/

ADD . /app

WORKDIR /app

RUN go build -o main .

FROM alpine:latest

COPY --from=builder /app/main .

CMD [ "./main" ]