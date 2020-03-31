all:
	go test -cover -coverprofile coverage.out
	go tool cover -func=coverage.out

deploy:
	git release
	git push --tags

index:
	GOPROXY=https://proxy.golang.org GO111MODULE=on go get github.com/jpedro/color@$(shell git tag | sort -V | tail -n 1)
