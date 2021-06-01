.PHONY: build
build:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
	go build -a -x -o ./bin/ecutil-windows-amd64.exe ./ecutil.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -a -x -o ./bin/ecutil-linux-amd64 ./ecutil.go
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 \
	go build -a -x -o ./bin/ecutil-darwin-amd64 ./ecutil.go
