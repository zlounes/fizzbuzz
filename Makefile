.PHONY: build
build:
	go build -o fizzbuzz  main.go

serverHost=localhost
serverPort=8080
int1=3
int2=5
limit=40
string1=fizz
string2=buzz

.PHONY: clean
clean:
	rm -f fizzbuzz

.PHONY: run
runServer: build
	./fizzbuzz -port $(serverPort)

.PHONY: runClient
runClient:
	curl -v -d "int1=$(int1)&int2=$(int2)&limit=$(limit)&string1=$(string1)&string2=$(string2)" -X POST http://localhost:8080/fizzbuzz

.PHONY: runTests
runTests: 
	go test -v ./...

.PHONY: buildImage
buildImage:
	docker build -t fizzbuzz:1.0 -f Dockerfile .

.PHONY: runIT
runIT:
	docker-compose down
	docker-compose up -d
	curl -v -d "int1=$(int1)&int2=$(int2)&limit=$(limit)&string1=$(string1)&string2=$(string2)" -X POST http://localhost:8080/fizzbuzz
	docker-compose down

.PHONY: help
help:
	@echo 'Targets:'
	@echo '  build        - build fizzbuzz executable'
	@echo '  runServer    - launch server'
	@echo '  runClient    - call fizzbuzz server'
	@echo '  runTests     - run unit tests'
	@echo '  buildImage   - build docker image'
	@echo '  runIT        - IT test on docker image'
	@echo 'Options:'
	@echo '  serverPort  - server port'
	@echo '  int1        - int1 value'
	@echo '  int2        - int2 value'
	@echo '  limit       - limit value'
	@echo '  string1     - string1 value'
	@echo '  string2     - string2 value'

