fmt:
	gofmt -w .

test:
	go test ./...

lint:
	golint ./...