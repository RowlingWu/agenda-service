FROM golang:1.8
RUN mkdir -p /go/src/github.com/RowlingWu/agenda-service
ADD . /go/src/github.com/RowlingWu/agenda-service/
WORKDIR /go/src/github.com/RowlingWu/agenda-service/cli
RUN go install
WORKDIR /go/src/github.com/RowlingWu/agenda-service/service
RUN go install
WORKDIR /go
CMD ["service"]