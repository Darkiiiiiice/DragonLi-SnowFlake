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
# go install github/mariomang/catrouter
go install io/mariomang/github
go install pingtest

CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o snowflake-32.exe ./src/io/mariomang/github/main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o snowflake-64.exe ./src/io/mariomang/github/main.go

mv snowflake-32.exe ./bin/
mv snowflake-64.exe ./bin/
#test 
go test -v io/mariomang/github/snowflake
# go test -v -bench=".*" io/mariomang/github/snowflake
# echo 'go success'
# if [ -f "$execute" ]; then 
#      ./bin/github
# fi
