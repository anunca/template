# Symfony nginx
## overview
- [doc](#doc)
- [install](#install)
- [notes](#notes)
## doc
- https://symfony.com/doc/current/index.html
## install
```sh
make help
```
## notes
prod
```sh
export ENV=prod
```
builder
```sh
export BUILDER=composer
```
```sh
export BUILDER=symfony
```
template
- Composer
```sh
export TEMPLATE=symfony/skeleton
```
```sh
export TEMPLATE=symfony/website-skeleton
```
- Symfony
```sh
export TEMPLATE=
```
```sh
export TEMPLATE=--full
```
version
```sh
export VERSION=3.4
```
```sh
export VERSION=4.4
```
```sh
export VERSION=5.4
```
```sh
export VERSION=6.1
```
create project
```sh
make project
```
debug
```sh
composer require --dev symfony/profiler-pack
```
maker
```sh
composer require --dev symfony/maker-bundle
```
dependency for make:controller
```sh
composer require annotations
composer require twig
```