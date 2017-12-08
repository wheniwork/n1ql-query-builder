#!/usr/bin/env bash

set -e

mode=atomic
profile="coverage.txt"

for d in $(go list ./... | grep -v vendor); do
	f="$(echo ${d} | tr / -).cover"
    go test -race -coverprofile=${f} -covermode=${mode} ${d}
done

echo "mode: $mode" > ${profile}
grep -h -v "^mode:" *.cover >> ${profile}