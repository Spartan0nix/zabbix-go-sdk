#!/bin/bash

case $1 in
--test)
    go test
;;
*)
    go run  $(go list ./... | grep -v _test)
;;
esac