docker_group := boundedinfinity
docker_image := echo
docker_ver   := 1.0
docker_tag   := $(docker_group)/$(docker_image):$(docker_ver)

make_dir := $(abspath $(shell pwd))
app_dir := /app

GOPACKAGE := github.com/boundedinfinity/echo

export GOPATH := $(make_dir):$(make_dir)/src/$(GOPACKAGE)/vendor
export GO15VENDOREXPERIMENT := 1
export PATH := $(PATH):$(GOPATH)/bin

.PHONY: list docker-build docker-run docker-push go-install glide-install

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

docker-build:
	docker build --tag $(docker_tag) .

docker-bash:
	docker run -it --rm $(docker_tag) bash

docker-push:
	docker push $(docker_tag)

go-install:
	go install $(GOPACKAGE)/...

glide-install:
	cd $(GOPATH)/src/$(GOPACKAGE) && glide install

revel-run:
	revel run $(GOPACKAGE)
