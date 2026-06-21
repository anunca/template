ARG NODE_IMAGE=node:26.1.0-alpine

FROM $NODE_IMAGE AS base

WORKDIR /app

COPY ./src/package.json ./src/package-lock.json /app/

ENTRYPOINT [ "npm", "run" ]

FROM base AS dev

RUN npm i -D nodemon

COPY ./src /app/

CMD [ "start:dev" ]

FROM base AS prod

COPY --from=dev /app /app

RUN npm i --omit=dev

CMD [ "start" ]
