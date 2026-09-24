ARG PYTHON_IMAGE=python:3.14.7-alpine

FROM $PYTHON_IMAGE AS base

RUN pip install django

WORKDIR /app

EXPOSE 8000

FROM base AS dev

CMD [ "python", "manage.py" , "runserver", "0.0.0.0:8000" ]

FROM dev AS build

COPY ./src/requirements.txt requirements.txt

#base dependencies
RUN echo "uWSGI>=2.0.18,<2.1" >> requirements.txt\
  && apk add --no-cache --virtual .tmp gcc libc-dev linux-headers\
  && pip install -r requirements.txt\
  && apk del .tmp

COPY ./src .

FROM base AS prod

COPY --from=build /app .

COPY --from=build /usr/local/bin/uwsgi /usr/local/bin/
COPY ./etc/uwsgi.ini .

COPY ./sh/entrypoint.sh /usr/local/bin/entrypoint
RUN chmod +x /usr/local/bin/entrypoint

CMD [ "/usr/local/bin/entrypoint" ]
