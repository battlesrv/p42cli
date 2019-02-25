build:
	go install ./

linux:
	GOOS=linux GOARCH=amd64 go build ./
	tar zcf p42cli.linux.amd64.tar.gz p42cli

darwin:
	GOOS=darwin GOARCH=amd64 go build ./
	tar zcf p42cli.darwin.amd64.tar.gz p42cli

.PHONY: build, linux, darwin
