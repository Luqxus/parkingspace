build:
	@go build -o ./bin/spacedrive
	
run: build
	@./bin/spacedrive
	
test:
	@go test ./...