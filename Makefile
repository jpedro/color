all:
	go test -cover -coverprofile coverage.out
	go tool cover -func=coverage.out
