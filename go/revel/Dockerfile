ARG GO_VERSION=1.26.3
ARG GO_IMAGE=golang:${GO_VERSION}-alpine

FROM $GO_IMAGE

RUN apk add --no-cache git

RUN go install github.com/revel/cmd/revel@latest

ENV PATH="$PATH:$GOPATH/bin"

WORKDIR $GOPATH

EXPOSE 9000

ENTRYPOINT [ "revel" ]
CMD [ "run", "-a", "./src/app" ]
