ARG NODE_IMAGE=node:26.7.0-alpine
ARG NGINX_IMAGE=nginx:1.31.4-alpine

FROM $NODE_IMAGE AS dev

# hadolint ignore=DL3016
RUN npm install -g create-react-app

WORKDIR /app

ENTRYPOINT [ "npm" ]
CMD [ "start" ]

FROM dev AS build

COPY src .

# hadolint ignore=DL3016
RUN npm install --omit=dev\
  && npm run build

FROM $NGINX_IMAGE AS prod

COPY --from=build /app/build /usr/share/nginx/html
