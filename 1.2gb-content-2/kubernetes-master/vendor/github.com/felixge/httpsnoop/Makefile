.PHONY: ci generate clean

ci: clean generate
	go test -race -v ./...

generate:
	go generate .

clean:
	rm -rf *_generated*.go
# ID-1768294467-e776d467
