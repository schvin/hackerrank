FROM golang:1.8
MAINTAINER George Lewis <schvin@schvin.net>
ENV REFRESHED_AT 2017-08-11

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["app"]
