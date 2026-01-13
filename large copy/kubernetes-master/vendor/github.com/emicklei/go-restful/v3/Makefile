all: test

test:
	go vet .
	go test -cover -v .

ex:
	find ./examples -type f -name "*.go" | xargs -I {} go build -o /tmp/ignore {}
# ID-1768294487-f3106779
