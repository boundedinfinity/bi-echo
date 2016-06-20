makefile_dir := $(abspath $(shell pwd))

docker_group := boundedinfinity
docker_image := echo
docker_ver   := 1.0.0
docker_tag   := $(docker_group)/$(docker_image):$(docker_ver)
docker_src   := /app
docker_port  := 8080

.PHONY: list

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

bootstrap:
	make go-bootstrap
	make bower-bootstrap

go-bootstrap:
	glide install
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/elazarl/go-bindata-assetfs/...

bower-bootstrap:
	bash -l -c 'bower install'

bindata-assetfs-clean:
	rm -f $(makefile_dir)/bindata_assetfs.go

bindata-assetfs:
	go-bindata-assetfs html/

bindata-clean:
	rm -f $(makefile_dir)/bindata.go

bindata:
	go-bindata html/

echo-debug:
	go-bindata-assetfs -debug -ignore=\\.gitignore view/... static/...
	go build
	$(makefile_dir)/echo

echo-run:
	go generate
	go build
	$(makefile_dir)/echo
