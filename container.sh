#!/bin/sh

source ./project.sh

TARGET_HOST=$(echo ${CONTAINER_HOST}|sed 's-tcp://--g'|cut -f1 -d:)
REV=$(git describe --long --tags --match='v*' --dirty 2>/dev/null || echo dev)


podman build -f container/build.Containerfile -t ${USER}/${PROJECT}:$REV . ||exit 1
podman tag ${USER}/${PROJECT}:$REV ${USER}/${PROJECT}:dev-latest
