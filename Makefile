docker_group := boundedinfinity
docker_image := echo
docker_ver   := 1.0
docker_tag   := $(docker_group)/$(docker_image):$(docker_ver)
app_dir := /app

go_package := github.com/boundedinfinity/echo
glide_pkg ?= none

make_dir := $(abspath $(shell pwd))
export GOPATH := $(make_dir)
export GO15VENDOREXPERIMENT := 1
export PATH := $(GOPATH)/bin:$(PATH)

.PHONY: list docker-build docker-run docker-push go-install glide-install revel-run

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

bin-path:
	@echo $(PATH)

docker-tag:
	@echo $(docker_tag)

docker-build:
	docker build --tag $(docker_tag) .

docker-bash:
	docker run -it --rm $(docker_tag) bash

docker-push:
	docker push $(docker_tag)

bootstrap: glide-install go-install

clean:
	go clean
	rm -rf $(GOPATH)/bin
	rm -rf $(GOPATH)/pkg
	rm -rf $(GOPATH)/src/$(go_package)/vendor

go-path:
	@echo $(GOPATH)

go-package:
	@echo $(go_package)

go-install:
	go install ./...

beego-run:
	 cd $(GOPATH)/src/$(go_package) && bee run
