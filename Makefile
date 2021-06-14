.PHONY: all
all: test


.PHONY: test
test:
	go test -cover -coverprofile coverage.out
	go tool cover -func coverage.out
	go tool cover -html coverage.out -o coverage.html

.PHONY: deploy
deploy:
	git release
	git push --tags

.PHONY: index
index:
	GOPROXY=https://proxy.golang.org GO111MODULE=on go get github.com/jpedro/color@$(shell git tag | sort -V | tail -n 1)
