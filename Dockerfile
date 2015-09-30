FROM golang

RUN go get github.com/ilgooz/stack

RUN cd /go/src/github.com/ilgooz/stack && make build

ENTRYPOINT /go/src/github.com/ilgooz/stack/stack \
  --mongo mongodb://mongo:27017/stack

EXPOSE 3000
