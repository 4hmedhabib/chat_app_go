SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)), /bin/ash,/bin/bash)

# Chat
chat-run:
	go run chat/api/services/cap/main.go | go run chat/api/tooling/logfmt/main.go


# Modules support
deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-list:
	go list -m -u -mod=readonly all

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all

