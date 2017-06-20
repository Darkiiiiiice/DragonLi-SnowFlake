#!/bin/sh
GOPATH=$(pwd)
echo $GOPATH
export GOPATH
echo $GOPATH

#clean
execute="./bin/github"
if [ -f "$execute" ]; then
    rm ./bin/github
fi

#download package
go get github.com/bmizerany/pq
go get github.com/julienschmidt/httprouter

#build program
go install github/mariomang/catrouter
go install io/mariomang/github
go install pingtest

#test 
go test io/mariomang/github

echo 'go success'
if [ -f "$execute" ]; then 
     ./bin/github
fi
