#!/bin/env bash

set -x
mkdir bin
GOOS=linux GOARCH=amd64 go build -o bin/swrpg_calc ./src
chmod +x bin/swrpg_calc
go env -w CGO_ENABLED=1
GOOS=windows GOARCH=amd64 go build -o bin/swrpg_calc.exe ./src