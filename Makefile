
all: test build

test:
	go test -cover ./...
	
build:
	go build -o bin/poker
	
docker:
	docker build --tag poker:latest .
	
clean:
	rm bin/*