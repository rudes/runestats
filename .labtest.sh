#!/bin/bash

go test -cover -v ./... | grep -o '^coverage: [0-9]*.[0-9]*\%' | awk '{print $2}' | sed 's/%//' | tr '\n' ' ' | awk '{printf "coverage: %.1f%%\n", ($1+$2)/2}'
