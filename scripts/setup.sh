#!/usr/bin/env bash


# Fetch go dependencies
go get -v -t -d ./...

# Fetch code generators
go get github.com/google/wire/cmd/wire