ARG GO_IMAGE=golang:1.27.0-alpine
ARG ALPINE_IMAGE=alpine:3.24.1

FROM $GO_IMAGE AS dev

WORKDIR /app

EXPOSE 3000

# hadolint ignore=DL3062
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

COPY --from=build /app/tmp/main /app/main

USER app

EXPOSE 3000

CMD [ "/app/main" ]