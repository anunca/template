#!/usr/bin/env bash
set -e

python manage.py collectstatic --noinput

# uwsgi --socket 0.0.0.0:8000 --master --enable-threads --module app.wsgi --processes 4 --threads 2
uwsgi uwsgi.ini
