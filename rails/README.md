# Rails
## overview
- [doc](#doc)
- [install](#install)
- [notes](#notes)
## doc
- https://rubyonrails.org/
- https://guides.rubyonrails.org/
- https://guides.rubyonrails.org/getting_started.html#mvc-and-you-generating-a-model
## install
```sh
make help
```
## notes
prod
```sh
export ENV=prod
```
Rails commands
```sh
docker compose exec app sh -c 'rake --tasks'
```
```sh
docker compose exec app sh -c 'rails'
```
```sh
docker compose exec app sh -c 'rails routes'
```
create project
- default template (sqlite3)
```sh
make project
```
template
```sh
export TEMPLATE='sqlite3'
```
```sh
export TEMPLATE='mysql'
```
```sh
export TEMPLATE='postgresql'
```
```sh
export TEMPLATE='sqlite3 --api'
```
TODO
-  was used in older Rails versions (Rails 6 / Webpacker era).
    - --webpack=react
- In Rails 7+
    - --javascript=esbuild
    - (npm i react react-dom)
```sh
export TEMPLATE='postgresql --webpack=react'
```