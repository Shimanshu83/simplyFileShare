build: 
	@go build -0 bin/fs 

run: 
	@./bin/fs 

test:
	@go test ./... -v 

