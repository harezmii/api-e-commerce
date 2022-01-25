FROM golang:1.17-alpine as builder
ENV GO111MODULE=on

RUN mkdir /app
ADD . /app


WORKDIR /app
COPY .  /app

RUN go mod download
RUN go mod verify

ENV CGO_ENABLED=0
#CMD gunicorn --bind 0.0.0.0:$3500 wsgi
RUN  go build  -o ecommerceapi /app/cmd/main.go


FROM alpine:latest AS production
COPY --from=builder /app /application
WORKDIR /application
CMD ["./ecommerceapi"]