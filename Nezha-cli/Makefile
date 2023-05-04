BINARY_NAME=nez
build:
	rm -rf bin/
	mkdir bin/
	GOOS=linux go build -o bin/${BINARY_NAME} main.go
run:
	./bin/${BINARY_NAME} 

build_and_run: 
	make build 
	make run 

clean:
	go clean
	rm bin/${BINARY_NAME}