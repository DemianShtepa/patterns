fmt:
	go fmt ./...
lint:
	golangci-lint run
coverage:
	go test -coverprofile=c.out ./...
view-coverage:
	go tool cover -html="c.out"
test:
	go test -cover -race ./...
mockery:
	mockery