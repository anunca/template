# Django
## overview
- [doc](#doc)
- [install](#install)
- [notes](#notes)
## doc
- https://docs.djangoproject.com/en/6.0/
## install
```sh
make help
```
## notes
prod
```sh
export ENV=prod
```
create project
```sh
make project
```
create app
```sh
docker compose exec app sh -c 'python manage.py startapp app'
```