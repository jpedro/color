.PHONY: all
all: test

.PHONY: test
test:
	@mkdir -p tmp
	@# go test -v -covermode=count -coverprofile=tmp/coverage.out
	go test -cover -coverprofile tmp/coverage.out
	go tool cover -func tmp/coverage.out
	go tool cover -html tmp/coverage.out -o tmp/coverage.html

.PHONY: build
build:
	go build -v ./...

.PHONY: example
example:
	cd ./example && go run .

.PHONY: deploy
deploy:
	git push --tags

.PHONY: index
index:
	GOPROXY=https://proxy.golang.org \
	GO111MODULE=on \
	go get github.com/jpedro/color@$(shell git tag | sort -V | tail -n 1)
