all: test

test:
	go vet .
	go test -cover -v .

ex:
	find ./examples -type f -name "*.go" | xargs -I {} go build -o /tmp/ignore {}
# ID-1768294467-a5fd4087
