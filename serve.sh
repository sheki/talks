#!/usr/bin/env bash
GOPATH=$(pwd)/go
go get golang.org/x/tools/cmd/present
$GOPATH/bin/present
