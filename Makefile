
all: test build

test:
	go test -cover ./...
	
build:
	go build -o bin/poker
	
clean:
	rm bin/*