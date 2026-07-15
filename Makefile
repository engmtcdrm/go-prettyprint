.PHONY: build runexe run test testv

build:
	echo "Size before build:"; ls -la examples |grep example; ls -lh examples |grep examples; echo "\n\nSize after build:"; go build --ldflags "-s -w" -o examples/examples ./examples; ls -la examples |grep examples; ls -lh examples |grep examples

check-build-files:
	go list -f '{{.GoFiles}}' .

runexe:
	./example/examples

run:
	go run ./examples/examples.go

test:
	go test .

testv:
	go test -v .

testcover:
	@go test -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html && rm coverage.out

testcoverall:
	@go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html && rm coverage.out
