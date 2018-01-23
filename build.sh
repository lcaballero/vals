#!/bin/bash
set -e

testing() {
	go test ./...
}

install() {
	go install
}

fmt() {
	gofmt -w .
}

clean() {
	gen=$(find . -type f | grep -e ".*gen\.go$")
	for f in $gen; do
		rm "$f"
	done
}

gen() {
	go generate ./... && fmt
}

all() {
	fmt
	install
	testing
}

if [ "$1" == "" ]; then
	all
else
	$1 $*
fi
