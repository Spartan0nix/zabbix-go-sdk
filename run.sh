#!/bin/bash

case $1 in
--test)
    podman system renumber

    echo "Running container stack..."
    podman-compose -f ./build/container-compose-test.yml up -d

    TIMER=40
    echo "Waiting $TIMER\\s for Zabbix server to initialize"
    i=1
    while [[ $i -ne $TIMER ]]
    do
        echo "$i / $TIMER"
        sleep 1
        i=$((i+1))
    done

    echo "Running test..."
    go test -v

    echo "Destroying container stack"
    podman-compose -f ./build/container-compose-test.yml down
;;
--up)
    podman system renumber
    podman-compose -f ./build/container-compose.yml up
;;
--down)
    podman-compose -f ./build/container-compose.yml down
;;
*)
    go run  $(go list ./... | grep -v _test)
;;
esac