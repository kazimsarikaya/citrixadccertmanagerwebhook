#!/bin/sh

source ./project.sh

TARGET_HOST=$(echo ${CONTAINER_HOST}|sed 's-tcp://--g'|cut -f1 -d:)
REV=$(git describe --long --tags --match='v*' --dirty 2>/dev/null || echo dev)


docker build -f docker/build.Dockerfile -t ${USER}/${PROJECT}:$REV . ||exit 1
docker tag ${USER}/${PROJECT}:$REV ${USER}/${PROJECT}:dev-latest
