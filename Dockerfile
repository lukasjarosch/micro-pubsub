FROM golang

COPY . /go/src/github.com/lukasjarosch/pubsub
WORKDIR /go/src/github.com/lukasjarosch/pubsub

RUN go get ./...
RUN go build
CMD pubsub


