SHELL = /bin/bash

.PHONY: test coverage
up:
	docker compose -f ./build/docker-compose.yml up

log:
	docker compose -f ./build/docker-compose.yml logs

down:
	docker compose -f ./build/docker-compose.yml down

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