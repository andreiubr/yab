FROM golang:1.7-alpine

COPY . /go/src/github.com/andreiubr/yab

RUN apk update
RUN apk add git

RUN cd /go/src/github.com/andreiubr/yab && go get github.com/Masterminds/glide
RUN cd /go/src/github.com/andreiubr/yab && glide install
RUN cd /go/src/github.com/andreiubr/yab && go install .

ENTRYPOINT ["yab"]

