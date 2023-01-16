#!/bin/sh

source ./project.sh

TARGET_HOST=$(echo ${CONTAINER_HOST}|sed 's-tcp://--g'|cut -f1 -d:)
REV=$(git describe --long --tags --match='v*' --dirty 2>/dev/null || echo dev)

CONTAINERREPO=${CONTAINERREPO:-localhost}

podman build -f container/build.Containerfile -t ${CONTAINERREPO}/${REPOUSER}/${REPOPROJECT}:$REV . ||exit 1
podman tag ${CONTAINERREPO}/${REPOUSER}/${PROJECT}:$REV ${CONTAINERREPO}/${REPOUSER}/${REPOPROJECT}:dev-latest
