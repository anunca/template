ARG NODE_IMAGE=node:26.7.0-alpine

FROM $NODE_IMAGE AS dev

WORKDIR /app

ENTRYPOINT [ "npm" ]
CMD [ "start" ]

FROM dev AS prod

COPY src/package.json .

RUN npm install --omit=dev

COPY src .
