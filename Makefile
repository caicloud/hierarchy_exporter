# Copyright 2017 The Caicloud Authors.
#
# The old school Makefile, following are required targets. The Makefile is written
# to allow building multiple binaries. You are free to add more targets or change
# existing implementations, as long as the semantics are preserved.
#
#   make        - default to 'build' target
#   make lint   - code analysis
#   make test   - run unit test (or plus integration test)
#   make build        - alias to build-local target
#   make build-local  - build local binary targets
#   make build-linux  - build linux binary targets
#   make container    - build containers
#   $ docker login registry -u username -p xxxxx
#   make push    - push containers
#   make clean   - clean up targets
#
# Not included but recommended targets:
#   make e2e-test
#
# The makefile is also responsible to populate project version information.
#
# TODO: implement 'make push'

#
# Tweak the variables based on your project.
#
ROOT=github.com/caicloud/hierarchy_exporter
TARGET=hierarchy-exporter

# Current version of the project.
VERSION ?= v0.2.1
REGISTRIES ?= cargo.caicloudprivatetest.com/caicloud

.PHONY: build container push

build-linux:
	@for registry in $(REGISTRIES); do                                                \
		docker run --rm                                                                \
		  -v ${PWD}:/go/src/$(ROOT)                                                    \
		  -w /go/src/$(ROOT)                                                           \
		  -e GOOS=linux                                                                \
		  -e GOARCH=amd64                                                              \
		  -e CGO_ENABLED=0                                                             \
		  -e GOPATH=/go                                                                \
		  $${registry}/golang:1.9.2-alpine3.6                                        \
		  go build -i -v .;                               			       \
	done                                                                             \

container: build-linux
	@for registry in $(REGISTRIES); do                                              \
		docker build -t $${registry}/${TARGET}:${VERSION} .;			\
	done									        \

push: container
	@for registry in $(REGISTRIES); do                                              \
		docker push $${registry}/${TARGET}:${VERSION};				\
	done										\
