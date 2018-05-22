
FROM golang:1.10.1

RUN mkdir -p /go/src/github.com/NghiaTranUIT/artify-core
WORKDIR /go/src/github.com/NghiaTranUIT/artify-core

ADD . /go/src/github.com/NghiaTranUIT/artify-core

RUN go get -v
