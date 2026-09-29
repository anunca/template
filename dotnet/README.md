# .NET
## overview
- [doc](#doc)
- [install](#install)
- [notes](#notes)
## doc
- https://docs.microsoft.com/fr-fr/aspnet/core/?view=aspnetcore-5.0/
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
- default template (webapi)
```sh
make project
```
templates
```sh
export TEMPLATE=webapp
```
```sh
export TEMPLATE=mvc
```
```sh
export TEMPLATE=webapi
```
TODO nodejs required
```sh
export TEMPLATE=angular
```
```sh
export TEMPLATE=react
```
```sh
export TEMPLATE=reactredux
```

Production builds require a project in `src` (run `make project` first for a new app).
Build with `make ENV=prod build`. The production image serves HTTP on container
port 5000, exposed on host port 80. Configure HTTPS at your reverse proxy.
