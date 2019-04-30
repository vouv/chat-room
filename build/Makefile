# GOOS：darwin、freebsd、linux、windows
# GOARCH：386、amd64、arm

all: darwin

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/app cmd/app.go

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app cmd/main.go

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/app.exe cmd/main.go

clean:
	rm -rf ./bin

.PHONY:darwin
.PHONY:linux
.PHONY:windows
.PHONY:clean