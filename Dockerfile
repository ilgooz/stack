FROM golang

WORKDIR /go/src/github.com/ilgooz/stack

ADD . .

RUN go get ./... && make build

ENTRYPOINT ./stack \
  --mongo mongodb://mongo:27017/stack

EXPOSE 3000
