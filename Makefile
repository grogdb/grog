.PHONY: clean clean-all deps gazelle goimports golint govet package protos test update

ANTLR:=$(shell which antlr4)
BAZEL:=$(shell which bazel)
DEP:=$(shell which dep)
GAZELLE:=$(shell which gazelle)
GO:=$(shell which go)
GOIMPORTS:=$(shell which goimports)
GOLINT:=$(shell which golint)
PROTOC:=$(shell which protoc)
PROTO_SOURCES:=$(shell find . -type f -name '*.proto')
PROTO_FILES:=$(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))

all: package

antlr:
	@$(ANTLR) -Dlanguage=Go -package query internal/parser/query/Query.g4
	@$(ANTLR) -Dlanguage=Go -package schema internal/parser/schema/Schema.g4

clean:
	@$(BAZEL) clean

clean-all:
	@$(BAZEL) clean --expunge

deps:
	$(DEP) ensure

goimports:
	@$(GOIMPORTS) -w $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./bazel*")

gazelle:
	@$(BAZEL) run //:gazelle

golint:
	@$(GOLINT) -set_exit_status $(shell go list ./... | grep -v /vendor/)

govet:
	@$(GO) vet $(shell go list ./... | grep -v /vendor/)

protos: $(PROTO_FILES)

$(PROTO_FILES): %.pb.go: %.proto

%.pb.go:
	@echo Compiling $<
	@$(PROTOC) --go_out=plugins=grpc:. $<

package: update
	@$(BAZEL) run --cpu=k8 //:package

test: protos
	@$(GO) test $(shell go list ./... | grep -v /vendor/)

update: goimports gazelle


# get docker image targets
# bazel query 'filter("docker-image", kind(app_layer, //...))'
# bazel query 'kind(".*_layer", //...)'
# bazel run --cpu=k8 $(bazel query 'kind("app_layer", //...)')