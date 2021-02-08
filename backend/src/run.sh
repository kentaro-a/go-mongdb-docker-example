#!/bin/sh
$GOPATH/bin/reflex -r '(\.go$|go\.mod)' -s go run main.go
