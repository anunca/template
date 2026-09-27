ARG NODE_IMAGE=node:26.7.0-alpine

FROM $NODE_IMAGE

# hadolint ignore=DL3016
RUN npm install -g serverless

WORKDIR /app

ENTRYPOINT [ "sls" ]
CMD ["-s", "dev", "offline", "start", "--watch"]
