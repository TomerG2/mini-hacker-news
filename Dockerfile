# syntax=docker/dockerfile:1

FROM golang:1.18

# set work dir in order to enable go-acc to run from test.sh
WORKDIR /build/mini-hacker-news

# copy requirements
ADD ./go.mod /build/mini-hacker-news/
ADD ./go.sum /build/mini-hacker-news/

# build and install the source code
RUN go mod download

# copy source code
ADD . /build/mini-hacker-news/
RUN go build .

ENTRYPOINT ["/build/mini-hacker-news/mini-hacker-news"]