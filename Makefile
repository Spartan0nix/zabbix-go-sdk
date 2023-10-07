SHELL = /bin/bash

.PHONY: test coverage
# ------------------------------------------------
# Docker
# ------------------------------------------------
up:
	docker compose -f ./build/docker-compose.yml up

down:
	docker compose -f ./build/docker-compose.yml down

up-test:
	docker compose -f ./build/docker-compose-test.yml up

down-test:
	docker compose -f ./build/docker-compose-test.yml down

# ------------------------------------------------
# Test
# ------------------------------------------------
test:
	@echo "Running container stack..."
	docker compose -f ./build/docker-compose-test.yml up -d
	@TIMER=40; \
	echo "Waiting $$TIMER\\s for Zabbix server to initialize"; \
	i=1; \
	while [[ $$i -ne $$TIMER ]]; \
	do \
		echo "$$i / $$TIMER"; \
		sleep 1; \
		i=$$((i+1)); \
	done;
	@echo "Running test..."
	go test
	@echo "Destroying container stack"
	docker compose -f ./build/docker-compose-test.yml down

coverage:
	@echo "Running container stack..."
	docker compose -f ./build/docker-compose-test.yml up -d
	@TIMER=40; \
	echo "Waiting $$TIMER\\s for Zabbix server to initialize"; \
	i=1; \
	while [[ $$i -ne $$TIMER ]]; \
	do \
		echo "$$i / $$TIMER"; \
		sleep 1; \
		i=$$((i+1)); \
	done;
	@echo "Running test..."
	go test -coverprofile=coverage.out
	@echo "Formatting coverage report to HTML..."
	go tool cover -html=coverage.out -o=coverage.html
	@echo "Destroying container stack"
	docker compose -f ./build/docker-compose-test.yml down

# ------------------------------------------------
# Bench
# ------------------------------------------------
# -run=^$                   : prevent Test functions from running
# -bench ^Benchmark.+$      : run functions starting with Benchmark
# ./...                     : run bench for every packages and sub-packages
# -args                     : Since cpuprofile can't be used for multiple packages, 'args' allow to pass argument for each individual packages.
#                             Kind a kike : foreach(pkg in pkgs) -> go test pkg -arg -> output pkg/cpu.prof ; pkg/mem.prof
#     -test.cpuprofile cpu.prof : output bench stats to a 'cpu.prof' file
#     -test.benchmem            : print memory allocations for benchmarks
#     -test.memprofile mem.prof : output meme bench stats to a 'mem.prof' file
run-bench:
	@echo "Running benchmarks..."
	go test -run=^$$ -bench ^Benchmark.+$$ -benchmem -count 6 .

run-bench-regex:
	@echo "Running benchmarks...(with verbose output)"
	go test -run=^$$ \
        -bench "${REGEX}" \
        -v \
        -cpuprofile cpu.prof \
        -benchmem \
        -memprofile mem.prof  \
        -count 6

profile-cpu:
	go tool pprof cpu.prof

profile-mem:
	go tool pprof mem.prof

flame-cpu:
	go tool pprof -http="localhost:8081" cpu.prof

flame-mem:
	go tool pprof -http="localhost:8081" mem.prof

# ------------------------------------------------
# Utils
# ------------------------------------------------
run:
	go run test/main.go