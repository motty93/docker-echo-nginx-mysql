# golang(echo) nginx mysql for docker template
## environment
* docker version 19.03.6
* docker-compose version 1.25.4
* go version 1.14.6
* mysql version 5.7
* nginx latest

## architecture
https://github.com/golang-standards/project-layout

## Get Start
### build

```
$ docker-compose build
```

### start

```
$ docker-compose up -d
```

### open url

```
# backend api
$ open http://localhost:8020 # for osx

or

$ xdg-open http://localhost:8020 # for linux
```

## container
### download library

```
$ docker-compose run --rm golang go get <package path>
```
