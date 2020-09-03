FROM golang:alpine

RUN mkdir -p /app

WORKDIR /app

ADD . /app
RUN go build ./main.go
EXPOSE 9090
CMD ["./main"]