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

.PHONY: runCurl
runCurl:
	curl -v -d "int1=$(int1)&int2=$(int2)&limit=$(limit)&string1=$(string1)&string2=$(string2)" -X POST http://localhost:$(serverPort)/fizzbuzz

.PHONY: runTests
runTests: 
	go test -v ./...

.PHONY: buildImage
buildImage:
	docker build -t fizzbuzz:1.0 -f Dockerfile .

.PHONY: runIT
runIT: export SERVER_PORT = $(serverPort)
runIT: 
	docker-compose down
	docker-compose build
	docker-compose up -d
	go test -v ./server_test.go -tags=integration
	docker-compose down

.PHONY: help
help:
	@echo 'Targets:'
	@echo '  build        - build fizzbuzz executable'
	@echo '  runServer    - launch local server'
	@echo '  runCurl      - use curl to call local server'
	@echo '  runTests     - run unit tests'
	@echo '  buildImage   - build docker image'
	@echo '  runIT        - run integration test on the docker image'
	@echo 'Options:'
	@echo '  serverPort  - server port'
	@echo '  int1        - int1 value'
	@echo '  int2        - int2 value'
	@echo '  limit       - limit value'
	@echo '  string1     - string1 value'
	@echo '  string2     - string2 value'

