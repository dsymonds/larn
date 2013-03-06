#! /bin/bash -e

cd $(dirname $0)
ID=$(git rev-parse HEAD | cut -c1-7)
go build -v -ldflags "-X main.larnBuildID $ID"
echo 1>&2 "OK"
