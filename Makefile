.PHONY: dep test ncommit pull

VERSION = v0.0.101

init:
	go mod init github.com/nelsonsaake/go

dep:
	@echo $(VERSION) > VERSION.md
	go mod tidy
	git add .
	git commit -m "go: changes for $(VERSION)"
	git tag $(VERSION) @REM Create a version tag

	@REM Push the version tag
	git push origin $(VERSION)

	@REM  Delete the existing latest tag if it exists and create a new one
	@if git rev-parse "latest" >nul 2>&1; then \
		git tag -d latest; \
		git push origin :refs/tags/latest; \
	fi

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