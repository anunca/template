ARG NODE_IMAGE=node:26.7.0-alpine
ARG NGINX_IMAGE=nginx:1.31.4-alpine

FROM $NODE_IMAGE AS base

# hadolint ignore=DL3016
RUN npm install -g create-react-app

WORKDIR /app

FROM base AS dev

# COPY src .

ENTRYPOINT [ "npm" ]
CMD [ "start" ]

FROM base AS prod-builder

COPY src .

# hadolint ignore=DL3016
RUN npm install --omit=dev\
  && npm run build

FROM $NGINX_IMAGE AS prod

COPY --from=prod-builder /app/build /usr/share/nginx/html
