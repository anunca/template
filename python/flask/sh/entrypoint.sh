#!/usr/bin/env bash
set -e

# uwsgi --socket 0.0.0.0:5000 --master --enable-threads --wsgi-file app.py --callable app --processes 4 --threads 2
uwsgi uwsgi.ini
