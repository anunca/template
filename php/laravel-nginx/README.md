# Laravel nginx
## overview
- [doc](#doc)
- [install](#install)
- [notes](#notes)
## doc
- https://laravel.com/
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
.env
```sh
cat src/.env.example > src/.env\
&& docker compose exec app php artisan key:generate
```
debugbar
```sh
docker compose exec app composer require barryvdh/laravel-debugbar --dev
```
cache
```sh
docker compose exec app php artisan cache:clear\
&& docker compose exec app php artisan config:cache
```