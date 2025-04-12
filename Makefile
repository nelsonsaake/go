.PHONY: dep dep2 test ncommit pull dd
.SILENT:
.ONESHELL:

VERSION := $(strip $(subst VERSION=,,$(firstword $(filter VERSION=%, $(shell type .env)))))

init:
	go mod init github.com/nelsonsaake/go

v++:
	cd cmd/v++
	go run .

dep: v++
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
	cd tests
	go test -run TestEntries -v

ncommit:
	git add .
	git commit -m "ncommit"
	git push origin main

pull:
	git pull origin main 

latest = $(shell git tag -l latest)
dd:
	echo $(latest)

