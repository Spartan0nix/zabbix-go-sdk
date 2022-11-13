#!/bin/bash

case $1 in
--test)
    go test
;;
--up)
    podman system renumber
    podman-compose up
;;
--down)
    podman-compose down
;;
*)
    go run  $(go list ./... | grep -v _test)
;;
esac