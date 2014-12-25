SHELL = /bin/bash

lg2static = $(GOPATH)/pkg/linux_amd64/github.com/libgit2/git2go.a
git2go = $(GOPATH)/src/github.com/libgit2/git2go
lg2dir = $(GOPATH)/src/github.com/libgit2/git2go/vendor
bsdir = $(GOPATH)/src/github.com/gaffo/boxstrapper

all: $(lg2static)
	cd $(bsdir)
	go get -u -t
	go generate
	go fmt
	go test github.com/gaffo/boxstrapper
	go install github.com/gaffo/boxstrapper

$(lg2static): $(lg2dir)
	@echo lg2static
	cd $(git2go) && make install

$(git2go):
	@echo git2go
	mkdir -p $(GOPATH)/src/github.com/libgit2
	cd $(GOPATH)/src/github.com/libgit2 && git clone https://github.com/libgit2/git2go

$(lg2dir): $(git2go)
	@echo lg2dir
	cd $(git2go) && git submodule update --init
