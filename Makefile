all:
	go test -cover -coverprofile coverage.out
	go tool cover -func=coverage.out

deploy:
	git release
	git push --tags
