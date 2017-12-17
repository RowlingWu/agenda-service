FROM golang:1.8
RUN mkdir -p /go/src/github.com/RowlingWu/agenda-test
ADD . /go/src/github.com/RowlingWu/agenda-test/
WORKDIR /go/src/github.com/RowlingWu/agenda-test/cli
RUN go get -d -v
RUN go install
WORKDIR /go/src/github.com/RowlingWu/agenda-test/service
RUN go get -d -v
RUN go install
WORKDIR /go
CMD ["service"]