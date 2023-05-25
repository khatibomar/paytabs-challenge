## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
.PHONY: confirm 
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## go/deps-reset: reset all the dependencies in the project
.PHONY: go/deps-reset
go/deps-reset: confirm
	git checkout -- go.mod
	go mod tidy
	go mod vendor
## go/tidy: get all new dependencies and vendor them 
.PHONY: go/tidy
go/tidy: confirm
	go mod tidy
	go mod vendor
## go/deps-upgrade: upgrade current dependencies
.PHONY: go/deps-upgrade
go/deps-upgrade: confirm
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

## go/deps-cleancache: clean mod cache
.PHONY: go/deps-cleancache
go/deps-cleancache: confirm
	go clean -modcache

## go/list: list all dependencies in the project
.PHONY: go/list
go/list:
	go list -mod=mod all

## test: will run all the tests
.PHONY: test
test:
	go test -race ./...

## run/v1: will run the server v1
.PHONY: run/v1
run/v1:
	go run ./...