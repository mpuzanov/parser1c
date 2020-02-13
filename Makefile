SOURCE=./cmd/parser1c

GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)	

.DEFAULT_GOAL = build 

build:
	go build -v ${SOURCE}

test: 
	go test -v  $(GO_TEST_DIRS)

run:
	go run ${SOURCE} -format xlsx ./example/kl_to.txt

lint :
	@goimports -w ${GO_SRC_DIRS}
	golangci-lint run
	
.PHONY: build test run lint