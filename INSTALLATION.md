# Hand's on go tooling setup

## Prerequisite

You must have `git` and `docker` installed on your machine.

## Manual installation

* [Install of Visual Studio Code](https://code.visualstudio.com/),
* [Install of Go Lang plugin in VSCode](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go),
* [Install Golang](https://golang.org/doc/install).
* Pull needed docker images

```
docker pull golang:1.9-alpine
docker pull mongo:3.4
```

### Workspace installation

* Create GOPATH tree structure,

```
.
├── bin
├── pkg
└── src
    └── github.com
        └── Sfeir
```

* Download go tools for Visual Studio Code,

```
go get -u -v \
    github.com/ramya-rao-a/go-outline \
    github.com/acroca/go-symbols \
    github.com/nsf/gocode \
    github.com/rogpeppe/godef \
    golang.org/x/tools/cmd/godoc \
    github.com/zmb3/gogetdoc \
    github.com/golang/lint/golint \
    github.com/fatih/gomodifytags \
    github.com/uudashr/gopkgs/cmd/gopkgs \
    golang.org/x/tools/cmd/gorename \
    sourcegraph.com/sqs/goreturns \
    github.com/cweill/gotests/... \
    golang.org/x/tools/cmd/guru \
    github.com/josharian/impl
```

* Download the git repository of hand's on go and checkout the `start` branch.

```
git clone -b start git@github.com:Sfeir/handsongo.git workspace/src/github.com/Sfeir/handsongo
```

### as script for linux/macos users

```
#!/bin/bash

mkdir -p workspace/src/github.com/Sfeir
mkdir -p workspace/bin
mkdir -p workspace/pkg

export GOPATH=$(pwd)/workspace

go get -u -v \
    github.com/ramya-rao-a/go-outline \
    github.com/acroca/go-symbols \
    github.com/nsf/gocode \
    github.com/rogpeppe/godef \
    golang.org/x/tools/cmd/godoc \
    github.com/zmb3/gogetdoc \
    github.com/golang/lint/golint \
    github.com/fatih/gomodifytags \
    github.com/uudashr/gopkgs/cmd/gopkgs \
    golang.org/x/tools/cmd/gorename \
    sourcegraph.com/sqs/goreturns \
    github.com/cweill/gotests/... \
    golang.org/x/tools/cmd/guru \
    github.com/josharian/impl

git clone -b start git@github.com:Sfeir/handsongo.git workspace/src/github.com/Sfeir/handsongo
```

## Offline package

* Take the package `handsongo-<OS>.zip` and unzip it,
* Go in this folder and read the `README.md` file.

Package available for `linux`, and `macos`.

### Warning for Windows users

Please use a Linux VM with virtual box instead.

## Need Docker images?

* Take the package `handsongo-docker.zip` and unzip it,
* Go in this folder and read the `README.md` file.
