#!/usr/bin/env bash

set -e

profile="coverage.txt"

for d in $(go list ./... | grep -v vendor); do
	f="$(echo ${d} | tr / -).cover"
    go test -race -coverprofile=${d} -covermode=atomic ${d}
done

echo "mode: $mode" > ${profile}
grep -h -v "^mode:" *.cover >> ${profile}