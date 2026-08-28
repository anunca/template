ARG NODE_IMAGE=node:26.7.0-alpine

FROM $NODE_IMAGE

RUN npm install -g serverless

WORKDIR /app

ENTRYPOINT [ "sls" ]
CMD ["-s", "dev", "offline", "start", "--watch"]
