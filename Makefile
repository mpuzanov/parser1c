.PHONY: build
build:
	go build -v .

test:
	go test

run:
	go build -o parser1c.exe && ./parser1c.exe -format xlsx ./files/kl_to.txt