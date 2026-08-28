# Express
## overview
- [doc](#doc)
- [install](#install)
- [notes](#notes)
## doc
- https://expressjs.com/
## install
```sh
make help
```
## notes
install express
```sh
make run
```
- generator
```sh
npm install -g express-generator@4
```
- create project
```sh
express .
```
create dev project
```sh
make build start
```
create tag project
```sh
ENV=prod TAG=${ENV} make build start
```