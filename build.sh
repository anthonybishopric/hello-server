#!/usr/bin/env bash
set -ve
go build
mkdir -p target/bin
cp hello-server target/bin/launch
SHA=$(git rev-parse HEAD)
tar -cf hello_$SHA.tar.gz target/*