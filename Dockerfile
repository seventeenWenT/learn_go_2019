FROM golang:latest

MAINTAINER seveteenwen "seventeenwen@gmail.com"

WORKDIR $GOPATH/src/cobra


COPY . $GOPATH/src/cobra


RUN go build

EXPOSE 8888

ENTRYPOINT ["./cobra"]
