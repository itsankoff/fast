default: build

build:
	@go build -o ./bin/fast *.go

run: build
	@./bin/fast
