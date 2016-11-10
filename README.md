#HANDSONGO
[![Build Status](https://travis-ci.org/Sfeir/handsongo.svg?branch=master)](https://travis-ci.org/Sfeir/handsongo)
[![GoDoc](https://godoc.org/github.com/sebastienfr/handsongo?status.svg)](https://godoc.org/github.com/Sfeir/handsongo)
[![codebeat badge](https://codebeat.co/badges/e8b35982-ddbf-41e7-a975-ff79f3c69502)](https://codebeat.co/projects/github-com-sfeir-handsongo)
[![Software License](http://img.shields.io/badge/license-APACHE2-blue.svg)](https://github.com/Sfeir/handsongo/blob/master/LICENSE)

This project is meant to help learning go. It provides a basic implementation of a REST microservice exposing a CRUD API.
Data are persisted in a MongoDB NoSQL database and the application is deployed in Docker.

## Technical stack

* [Docker](https://www.docker.com)
* [MongoDB NoSQL database](https://www.mongodb.com)
* [Go is the language](https://golang.org)
* [Gorilla Mux the URL router](https://github.com/gorilla/mux)
* [Gorilla Mux the request context manager](https://github.com/gorilla/context)
* [Urfave negroni Web HTTP middleware](https://github.com/urfave/negroni)
* [Urfave cli the command line client parser](https://github.com/urfave/cli)
* [Sirupsen the logger](https://github.com/Sirupsen/logrus)
* [The database driver](https://gopkg.in/mgo.v2)
* [Godep the dependency manager](https://github.com/tools/godep)
* [Golint the source linter](https://github.com/golang/lint)

## Architecture

![main architecture](doc/img/main_architecture.png "Main architecture")

![web architecture](doc/img/web_architecture.png "Web architecture")

## Build

```shell
make help
```
