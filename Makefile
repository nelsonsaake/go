.PHONY: dep test ncommit

VERSION = v0.0.3

dep:
	go mod tidy
	git add .
	git commit -m "go: changes for $(VERSION)"
	git tag $(VERSION)
	git push origin main $(VERSION)
	@REM GOPROXY=proxy.golang.org go list -m https://github.com/nelsonsaake/go.git/mymodule@v0.1.0

test:
	go test ./...

ncommit:
	git add .
	git commit -m "ncommit"
	git push origin main
