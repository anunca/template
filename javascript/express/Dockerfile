ARG NODE_IMAGE=node:26.7.0-alpine

FROM $NODE_IMAGE AS base

WORKDIR /app

COPY ./src/package.json ./src/package-lock.json /app/

ENTRYPOINT [ "npm", "run" ]

FROM base AS dev

RUN npm install -D nodemon

COPY ./src /app/

CMD [ "start:dev" ]

FROM base AS prod

COPY --from=dev /app /app

RUN npm install --omit=dev

CMD [ "start" ]
