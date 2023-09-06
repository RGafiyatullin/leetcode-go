#!/bin/sh

go test $(
    find . -iname 'go.mod' \
    | while read f; do 
        dirname $f
    done)