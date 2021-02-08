#!/bin/sh
go generate 
$GOPATH/bin/reflex -r '(\.go$|go\.mod)' -s go run main.go
