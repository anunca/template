ARG NODE_IMAGE=node:26.1.0-alpine

FROM $NODE_IMAGE

RUN npm -i -g serverless

WORKDIR /app

ENTRYPOINT [ "sls" ]
CMD ["-s", "dev", "offline", "start", "--watch"]
