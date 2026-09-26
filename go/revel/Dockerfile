ARG GO_IMAGE=golang:1.27.0-alpine

FROM $GO_IMAGE

RUN apk add --no-cache --virtual .virtual git\
  && apk del .virtual

RUN go install github.com/revel/cmd/revel@latest

ENV PATH="$PATH:$GOPATH/bin"

WORKDIR $GOPATH

EXPOSE 9000

ENTRYPOINT [ "revel" ]
CMD [ "run", "-a", "./src/app" ]
