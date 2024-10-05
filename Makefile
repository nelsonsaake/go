.PHONY: dep test ncommit pull dd

VERSION = v0.0.122

init:
	go mod init github.com/nelsonsaake/go

dep:
	echo $(VERSION) > VERSION.md
	go mod tidy

	git add .
	git commit -m "go: changes for $(VERSION)"
	
	git tag $(VERSION) 
	git push origin $(VERSION)

	git tag -d latest
	git push origin :refs/tags/latest  

	git tag latest
	git push origin latest

	git push origin main

	echo deploy completed

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