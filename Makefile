makefile_dir 	:= $(abspath $(shell pwd))

docker_service	:= bi-echo
docker_tag		:= $(shell cat $(makefile_dir)/docker-compose.yml | shyaml get-value services.$(docker_service).image)
docker_src	 	:= /src

.PHONY: list

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

bootstrap:
	@make bower-bootstrap
	@make go-bootstrap

purge:
	@make bower-purge
	@make go-purge

clean:
	@make go-clean

docker-tag:
	@echo $(docker_tag)

docker-up:
	docker-compose up $(docker_service)

docker-stop:
	docker-compose stop $(docker_service)

docker-build:
	docker-compose build $(docker_service)

docker-bash:
	docker-compose run --rm $(docker_service) bash

docker-dev:
	docker-compose -f $(makefile_dir)/docker-compose-dev.yml run --rm $(docker_service) bash

docker-push:
	docker push $(docker_tag)

go-bootstrap:
	glide install

go-clean:
	go clean

go-purge:
	@make go-clean
	rm -rf $(makefile_dir)/vendor
	rm -rf $(makefile_dir)/glide.lock

bower-bootstrap:
	bower install

bower-purge:
	rm -rf $(makefile_dir)/$(shell cat $(makefile_dir)/.bowerrc | jq -r .directory)
