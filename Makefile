
sourse=./cmd/parser1c

.PHONY: build test run
build:
	go build -v ${sourse}

test:
	go test -v

run:
	go run ${sourse} -format xlsx ./files/kl_to.txt