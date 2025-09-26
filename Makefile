include .env

PREFIX=github.com/ajaincertinia/anlogtemp/scanner
LDFLAGS=-s -w -X '$(PREFIX).Version=$(VERSION)'

build:
	go build -ldflags "$(LDFLAGS)" .

install:
	go install -ldflags "$(LDFLAGS)" .


build-binaries:
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o anlogtemp.exe
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o anlogtemp_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o anlogtemp_darwin_arm64
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o anlogtemp_linux_amd64
	GOOS=linux GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o anlogtemp_linux_arm64

