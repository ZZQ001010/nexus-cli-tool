#!/usr/bin/env bash

PRJ=`pwd`
cd $PRJ/src
echo $PRJls
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build   -o $PRJ/bin/nexus-cli-linux $PRJ/src/main/main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build   -o $PRJ/bin/nexus-cli-windows $PRJ/src/main/main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -o $PRJ/bin/nexus-cli-mac $PRJ/src/main/main.go
chmod 777 $PRJ/bin/