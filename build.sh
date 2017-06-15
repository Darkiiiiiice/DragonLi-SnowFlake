#!/bin/sh
GOPATH=$(pwd)
echo $GOPATH
export GOPATH
echo $GOPATH
go install io/mariomang/github
echo 'go success'
./bin/github