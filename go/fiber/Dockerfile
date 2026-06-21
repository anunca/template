ARG GO_VERSION=1.26.3
ARG GO_IMAGE=golang:${GO_VERSION}-alpine
ARG ALPINE_VERSION=3.23.4
ARG ALPINE_IMAGE=alpine:${ALPINE_VERSION}

FROM $GO_IMAGE AS dev

WORKDIR /app

EXPOSE 3000

RUN go install github.com/air-verse/air@latest

CMD [ "air", "-c", ".air.toml" ]

FROM dev AS build

COPY ./src .

RUN go mod tidy && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main

FROM $ALPINE_IMAGE AS prod

RUN addgroup -S app \
&& adduser -S app -G app

COPY --from=build /app/main /app/main

USER app

EXPOSE 3000

CMD [ "/app/main" ]