
.PHONY: all clean linux/amd64 darwin/amd64 windows/amd64

all: 
	@mkdir -p ./bin/linux/amd64 ./bin/darwin/amd64 ./bin/windows/amd64
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux/amd64/github-importer-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/amd64/github-importer-darwin-amd64
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows/amd64/github-importer-windows-amd64.exe

clean:
	@rm -rf ./bin
