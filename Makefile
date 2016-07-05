makefile_dir := $(abspath $(shell pwd))

docker_group := boundedinfinity
docker_image := echo
docker_ver   := 1.0.0
docker_tag   := $(docker_group)/$(docker_image):$(docker_ver)
docker_src   := /app
docker_port  := 8080

go_assetfile := $(makefile_dir)/bindata_assetfs.go

.PHONY: list

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

bootstrap:
	make go-bootstrap
	make bower-bootstrap

go-bootstrap:
	glide install
	go get github.com/githubnemo/CompileDaemon
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/elazarl/go-bindata-assetfs/...

bower-bootstrap:
	bash -l -c 'bower install'

bindata-assetfs-clean:
	rm -f $(makefile_dir)/bindata_assetfs.go

bindata-clean:
	rm -f $(go_assetfile)

echo-build:
	go generate
	go build

echo-debug:
	make echo-build
	$(makefile_dir)/echo

echo-run:
	go generate
	go build
	$(makefile_dir)/echo

echo-refresh:
	CompileDaemon \
		-build="make echo-build" \
		-command $(makefile_dir)/echo \
		-directory=$(makefile_dir) -color \
		-exclude-dir=.git -exclude-dir=.idea -exclude-dir=static -exclude-dir=vendor \
		-exclude=bindata_assetfs.go -exclude=bindata.go -exclude=echo
