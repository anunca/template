ARG PYTHON_IMAGE=python:3.14.7-alpine

FROM $PYTHON_IMAGE AS base

# hadolint ignore=DL3013
RUN pip install --no-cache-dir flask

WORKDIR /app

EXPOSE 5000

FROM base AS dev

CMD [ "python", "-m" , "flask", "run", "--host=0.0.0.0" ]

FROM dev AS build

COPY ./src/requirements.txt requirements.txt

#base dependencies
RUN echo "uWSGI>=2.0.18,<2.1" >> requirements.txt\
  && apk add --no-cache --virtual .build-deps gcc libc-dev linux-headers\
  && pip install --no-cache-dir -r requirements.txt\
  && apk del .build-deps

COPY ./src .

FROM base AS prod

COPY --from=build /app .

COPY --from=build /usr/local/bin/uwsgi /usr/local/bin/
COPY ./etc/uwsgi.ini .

COPY ./sh/entrypoint.sh /usr/local/bin/entrypoint
RUN chmod +x /usr/local/bin/entrypoint

CMD [ "/usr/local/bin/entrypoint" ]
