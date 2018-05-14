#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

OS_DIST_NAME=$1
OS_NAME=$2
OS_ARCH=$3
GOLANG_VERSION=$4
GOLANG_DEP_VERSION=$5
VSCODE_OS_VERSION=$6
VSCODE_GOLANG_PLUGIN_VERSION=$7
JQ_VERSION=$8
JQ_OS_ARCH=$9

SCRIPT_SCOPE="${OS_DIST_NAME} distribution"

printf "%-20s | %-40s" "$SCRIPT_SCOPE" "prepare for ${OS_NAME} ${OS_ARCH}"
rm -rf "dist/${OS_DIST_NAME}"
mkdir -p "dist/${OS_DIST_NAME}/apps"
cp "${OS_DIST_NAME}/"* "dist/${OS_DIST_NAME}/"
printf "\e[1;32m%s\e[0m\n" "DONE"



printf "%-20s | %-40s" "$SCRIPT_SCOPE" "prepare workspace"
mkdir -p "dist/${OS_DIST_NAME}/workspace/src/github.com/Sfeir"
mkdir -p "dist/${OS_DIST_NAME}/workspace/bin"
mkdir -p "dist/${OS_DIST_NAME}/workspace/pkg"
git clone -q -b start https://github.com/Sfeir/handsongo.git "dist/${OS_DIST_NAME}/workspace/src/github.com/Sfeir/handsongo"
printf "\e[1;32m%s\e[0m\n" "DONE"



printf "%-20s | %-40s" "$SCRIPT_SCOPE" "prepare golang (${GOLANG_VERSION})"
wget -q --output-document="$(pwd)/dist/${OS_DIST_NAME}/apps/go.tar.gz" "https://storage.googleapis.com/golang/go${GOLANG_VERSION}.${OS_NAME}-${OS_ARCH}.tar.gz"
tar -C "$(pwd)/dist/${OS_DIST_NAME}/apps/" -xzf "$(pwd)/dist/${OS_DIST_NAME}/apps/go.tar.gz"
rm "$(pwd)/dist/${OS_DIST_NAME}/apps/go.tar.gz"
printf "\e[1;32m%s\e[0m\n" "DONE"



printf "%-20s | %-40s" "$SCRIPT_SCOPE" "prepare dep (${GOLANG_DEP_VERSION})"
mkdir -p "$(pwd)/dist/${OS_DIST_NAME}/tmp_dep"
wget -q --output-document="$(pwd)/dist/${OS_DIST_NAME}/tmp_dep/dep" "https://github.com/golang/dep/releases/download/${GOLANG_DEP_VERSION}/dep-${OS_NAME}-${OS_ARCH}"
mkdir -p "$(pwd)/dist/${OS_DIST_NAME}/workspace/bin/"
mv "$(pwd)/dist/${OS_DIST_NAME}/tmp_dep/dep" "$(pwd)/dist/${OS_DIST_NAME}/workspace/bin/dep"
rm -rf "$(pwd)/dist/${OS_DIST_NAME}/tmp_dep"
printf "\e[1;32m%s\e[0m\n" "DONE"



printf "%-20s | %-40s" "$SCRIPT_SCOPE" "prepare vscode (stable)"
if [ "${OS_DIST_NAME}" == "macos" ]; then
    wget -q --output-document="$(pwd)/dist/${OS_DIST_NAME}/apps/vscode.zip" "https://go.microsoft.com/fwlink/?LinkID=${VSCODE_OS_VERSION}"
    unzip -q "$(pwd)/dist/${OS_DIST_NAME}/apps/vscode.zip" -d "$(pwd)/dist/${OS_DIST_NAME}/apps/"
    rm "$(pwd)/dist/${OS_DIST_NAME}/apps/vscode.zip"
fi

if [ "${OS_DIST_NAME}" == "linux" ]; then
    wget -q --output-document="$(pwd)/dist/${OS_DIST_NAME}/apps/vscode.tar.gz" "https://go.microsoft.com/fwlink/?LinkID=${VSCODE_OS_VERSION}"
    tar -C "$(pwd)/dist/${OS_DIST_NAME}/apps/" -xzf "$(pwd)/dist/${OS_DIST_NAME}/apps/vscode.tar.gz"
    rm "$(pwd)/dist/${OS_DIST_NAME}/apps/vscode.tar.gz"
fi
cp vscode_settings.json "dist/${OS_DIST_NAME}/"
printf "\e[1;32m%s\e[0m\n" "DONE"



printf "%-20s | %-40s" "$SCRIPT_SCOPE" "prepare vscode golang plugin (${VSCODE_GOLANG_PLUGIN_VERSION})"
wget -q --output-document="$(pwd)/dist/${OS_DIST_NAME}/apps/go.vsix" "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/ms-vscode/vsextensions/Go/${VSCODE_GOLANG_PLUGIN_VERSION}/vspackage"
printf "\e[1;32m%s\e[0m\n" "DONE"



printf "%-20s | %-40s" "$SCRIPT_SCOPE" "prepare vscode golang tools"
GOPATH="$(pwd)/dist/${OS_DIST_NAME}/workspace"
export GOPATH

GOOS="${OS_NAME}" GOARCH="${OS_ARCH}" go get -u \
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

if [ -d "dist/${OS_DIST_NAME}/workspace/bin/${OS_NAME}_${OS_ARCH}" ]; then
    mv "dist/${OS_DIST_NAME}/workspace/bin/${OS_NAME}_${OS_ARCH}/"* "dist/${OS_DIST_NAME}/workspace/bin/"
fi
if [ -d "dist/${OS_DIST_NAME}/workspace/bin/${OS_NAME}_${OS_ARCH}" ]; then
    rm -rf "dist/${OS_DIST_NAME}/workspace/bin/${OS_NAME}_${OS_ARCH}"
fi
printf "\e[1;32m%s\e[0m\n" "DONE"



printf "%-20s | %-40s" "$SCRIPT_SCOPE" "prepare jq (${JQ_VERSION})"
mkdir -p "$(pwd)/dist/${OS_DIST_NAME}/apps/bin"
wget -q --output-document="$(pwd)/dist/${OS_DIST_NAME}/apps/bin/jq" "https://github.com/stedolan/jq/releases/download/jq-${JQ_VERSION}/jq-${JQ_OS_ARCH}"
chmod +x "$(pwd)/dist/${OS_DIST_NAME}/apps/bin/jq"
printf "\e[1;32m%s\e[0m\n" "DONE"



printf "%-20s | %-40s" "$SCRIPT_SCOPE" "prepare archive"
cd "dist/${OS_DIST_NAME}" || exit
zip -q -r "../handsongo-${OS_DIST_NAME}.zip" ./*
printf "\e[1;32m%s\e[0m\n" "DONE"
