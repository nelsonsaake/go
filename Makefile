.PHONY: dep test ncommit pull dd

VERSION = v0.0.102

init:
	go mod init github.com/nelsonsaake/go

dep:
	@echo $(VERSION) > VERSION.md
	go mod tidy
	
	git add .
	git commit -m "go: changes for $(VERSION)"
	
	@REM Create a version tag
	git tag $(VERSION) 

	@REM Push the version tag
	git push origin $(VERSION)

	git tag -d latest;  
	git push origin :refs/tags/latest;  

	git tag latest # Tag the latest commit
	git push origin latest

	@REM GOPROXY=proxy.golang.org go list -m https://github.com/nelsonsaake/go.git@$(VERSION) 

	echo "--- DEPLOY COMPLETED ---"

test:
	go test ./...

ncommit:
	git add .
	git commit -m "ncommit"
	git push origin main

pull:
	git pull origin main 

latest = $(shell git tag -l latest)
dd:
	echo $(latest)