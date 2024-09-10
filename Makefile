build:
	@go build -o bin/ENS

run:build
	@./bin/ENS

test:
	@go test -v ./...

generateReport:
			  @./bin/coverageReport