FROM golang:1.17-alpine as builder
ENV GO111MODULE=on

RUN mkdir /app
ADD . /app


WORKDIR /app
COPY .  /app

RUN go mod download
RUN go mod verify


RUN go build -o ecommerceapi /app/cmd/main.go
CMD gunicorn --bind 0.0.0.0:$PORT wsgi

CMD ["./ecommerceapi"]