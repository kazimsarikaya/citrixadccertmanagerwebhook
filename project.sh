#!/bin/sh -x

CONTAINER_HOST="tcp://127.0.0.1:1928"
USER=kazimsarikaya
PROJECT=test
REPO=github.com/${USER}/${PROJECT}


if [[ "x${USER}" == "x" ]]; then
  echo set user
  exit 1
fi

if [[ "x${PROJECT}" == "x" ]]; then
  echo set project
  exit 1
fi

if [[ "x${REPO}" == "x" ]]; then
  echo set repo
  exit 1
fi

test "X$(basename -- "$0")" = "X.project.sh" || exit 0

cmd=${1:-build}

if [[ "x${cmd}" == "xinit" ]]; then
  echo Initializing project
  mkdir -p bin tmp cmd internal/${PROJECT} pkg/${PROJECT} tmp vendor
  go mod init ${REPO}
elif [[ "x${cmd}" == "xdestroy" ]]; then
  echo Destroying project
  rm -fr bin tmp cmd internal pkg tmp vendor go.mod go.sum *.test
fi
