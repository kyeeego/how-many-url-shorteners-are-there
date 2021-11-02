FROM golang:1.17.2-alpine3.14

RUN mkdir /app
ADD go.mod /app
ADD go.sum /app
WORKDIR /app

RUN go mod download

ADD . /app

RUN go build -o main .

CMD ["/app/main"]
