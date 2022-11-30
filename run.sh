#!/bin/bash

arg_1=$1
arg_2=$2

case $arg_1 in
--test)
    echo "Running container stack..."
    docker compose -f ./build/container-compose-test.yml up -d

    TIMER=40
    echo "Waiting $TIMER\\s for Zabbix server to initialize"
    i=1
    while [[ $i -ne $TIMER ]]
    do
        echo "$i / $TIMER"
        sleep 1
        i=$((i+1))
    done

    if [[ $arg_2 == "--coverage"* ]]
    then
        echo "Running test... (with coverage)"
        go test -coverprofile=coverage.out

        echo "Formatting coverage report to HTML..."
        go tool cover -html=coverage.out -o=coverage.html
    else
        echo "Running test..."
        go test
    fi

    echo "Destroying container stack"
    docker compose -f ./build/container-compose-test.yml down
;;
--up)
    docker compose -f ./build/container-compose.yml up
;;
--down)
    docker compose -f ./build/container-compose.yml down
;;
--run)
    go run  $(go list ./... | grep -v _test)
;;
*)
    echo "Unknown option(s) : $*"
    echo "Available options are :"
    echo "--test [--coverage]"
    echo "--up"
    echo "--down"
    echo "--run"
;;
esac