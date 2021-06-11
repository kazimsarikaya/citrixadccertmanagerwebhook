#!/bin/sh

source ./project.sh

TARGET_HOST=$(printenv CONTAINER_HOST|sed 's-tcp://--g'|cut -f1 -d:)
#rm -fr tmp/*

find bin -maxdepth 1 -name "*.test" -delete

for pkg in `go list ./...`;
do
  tf=${pkg#"${REPO}/"}
  tf=${tf//\//.}
  tf=bin/${tf}.test
  go test -cover -c -o $tf $pkg ||exit 1
done

if [ "x$1" == "xrun" ]; then
  for tf in $(cd bin;find . -type f -name "*.test");
  do
    ./bin/$tf -test.coverprofile tmp/${tf}.cover.out -ginkgo.v -test.v -v ${VERBOSE:-5} ||exit 1
    go tool cover -html tmp/${tf}.cover.out -o tmp/${tf}.cover.html
  done
fi
