#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'
SCRIPT_SCOPE="docker distribution"

DOCKER_GOLANG_IMAGE_VERSION=$1
DOCKER_MONGO_IMAGE_VERSION=$2

mkdir -p dist/docker

printf "%-20s | %-40s" "$SCRIPT_SCOPE" "get golang image (${DOCKER_GOLANG_IMAGE_VERSION})"
docker pull "golang:${DOCKER_GOLANG_IMAGE_VERSION}" > /dev/null
printf "\e[1;32m%s\e[0m\n" "DONE"

printf "%-20s | %-40s" "$SCRIPT_SCOPE" "get mongo image (${DOCKER_MONGO_IMAGE_VERSION})"
docker pull "mongo:${DOCKER_MONGO_IMAGE_VERSION}" > /dev/null
printf "\e[1;32m%s\e[0m\n" "DONE"

printf "%-20s | %-40s" "$SCRIPT_SCOPE" "save images"
docker save --output dist/docker/image-golang.tar "golang:${DOCKER_GOLANG_IMAGE_VERSION}"
docker save --output dist/docker/image-mongo.tar "mongo:${DOCKER_MONGO_IMAGE_VERSION}"
printf "\e[1;32m%s\e[0m\n" "DONE"

printf "%-20s | %-40s" "$SCRIPT_SCOPE" "prepare documentation"
cp docker/* dist/docker/
printf "\e[1;32m%s\e[0m\n" "DONE"

printf "%-20s | %-40s" "$SCRIPT_SCOPE" "prepare archive"
cd dist/docker || exit
zip -q -r ../handsongo-docker.zip ./*
printf "\e[1;32m%s\e[0m\n" "DONE"
