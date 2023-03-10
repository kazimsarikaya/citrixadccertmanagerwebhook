#!/bin/sh

source ./project.sh

cmd=${1:-build}

go mod tidy
go mod vendor

linterver=""
if [ -f bin/golangci-lint ]; then
  linterver=$(bin/golangci-lint version|gawk 'match($0, /version (.*) built/, a) {print a[1]}')
fi
lastlinerver=$(curl -s https://api.github.com/repos/golangci/golangci-lint/releases/latest|jq -r '.tag_name')

if [ "v$linterver" != "$lastlinerver" ]; then
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $lastlinerver
fi

bin/golangci-lint --timeout=10m run --fix  ||exit 1

for proto in $(find internal -name *.proto); do
  protoc --experimental_allow_proto3_optional -I $(dirname $proto) --go_out=$(dirname $proto) $(basename $proto)
done

if [ "x$cmd" == "xbuild" ]; then
  REV=$(git describe --long --tags --match='v*' --dirty 2>/dev/null || git rev-list -n1 HEAD)
  NOW=$(date +'%Y-%m-%d_%T')
  GOV=$(go version)
  go build -ldflags "${LDFLAGS} -X main.version=$REV -X main.buildTime=$NOW -X 'main.goVersion=${GOV}'"  -o ./bin/${REPOPROJECT} ./cmd
elif [ "x$cmd" == "xtest" ]; then
  shift
  ./test.sh $@
else
  echo unknown command
fi
