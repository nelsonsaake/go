.PHONY: dep test ncommit pull

VERSION = v0.0.39

init:
	go mod init github.com/nelsonsaake/go

dep:
	go mod tidy
	git add .
	git commit -m "go: changes for $(VERSION)"
	git tag $(VERSION)
	git push origin main $(VERSION)
	@REM GOPROXY=proxy.golang.org go list -m https://github.com/nelsonsaake/go.git@$(VERSION)

test:
	go test ./...

ncommit:
	git add .
	git commit -m "ncommit"
	git push origin main

pull:
	git pull origin main 