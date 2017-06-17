#!/bin/sh
GOPATH=$(pwd)
echo $GOPATH
export GOPATH
echo $GOPATH


execute="./bin/github"
if [ -f "$execute" ]; then
    rm ./bin/github
fi

go get github.com/bmizerany/pq
go get github.com/julienschmidt/httprouter
go install github/mariomang/catrouter
go install io/mariomang/github
echo 'go success'
if [ -f "$execute" ]; then 
     ./bin/github
fi