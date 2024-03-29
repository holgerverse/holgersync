FROM golang:1.18-alpine
LABEL holgerverse.holgersync.author="Holgerson97"

WORKDIR $GOPATH/src/github.com/holgerverse/holgersync

COPY . .

RUN go install

ENTRYPOINT [ "sh" ]