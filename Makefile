-include .env

VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE):$(GOBASE)/src/
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

FORMATTING_BEGIN_YELLOW = \033[0;33m
FORMATTING_BEGIN_BLUE = \033[36;m
FORMATTING_BEGIN_GREEN = \033[32;1m
FORMATTING_END = \033[0m
LOG := @printf -- "${FORMATTING_BEGIN_GREEN}%s ${FORMATTING_END}\n"

.PHONY: help
help: ## Show help message
		@clear
		@printf -- "${FORMATTING_BEGIN_BLUE}%s${FORMATTING_END}\n" \
		"|---------------------------------------------------|" \
        "|     ________       .__                            |" \
        "|    /  _____/  ____ |  | _____    ____    ____     |" \
        "|   /   \  ___ /  _ \|  | \__  \  /    \  / ___\    |" \
        "|   \    \_\  (  <_> )  |__/ __ \|   |  \/ /_/  >   |" \
        "|    \______  /\____/|____(____  /___|  /\___  /    |" \
        "|           \/                 \/     \//_____/     |" \
		"|---------------------------------------------------|" \
		""
		@awk 'BEGIN {\
				FS = ":.*##"; \
				printf                "Usage: ${FORMATTING_BEGIN_BLUE}OPTION${FORMATTING_END}=<value> make ${FORMATTING_BEGIN_YELLOW}<target>${FORMATTING_END}\n"\
		  } \
		  /^[a-zA-Z0-9_-]+:.*?##/ { printf "  ${FORMATTING_BEGIN_BLUE}%-46s${FORMATTING_END} %s\n", $$1, $$2 } \
		  /^.?.?##~/              { printf "   %-46s${FORMATTING_BEGIN_YELLOW}%-46s${FORMATTING_END}\n", "", substr($$1, 6) } \
		  /^##@/                  { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID := /tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

install: go-get ## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar

start: ## start: Start in development mode. Auto-starts when code changes.
	@bash -c "trap 'make stop' EXIT; $(MAKE) clean compile start-server watch run='make clean compile start-server'"

stop: stop-server ## stop: Stop development mode.

start-server: stop-server
	$(LOG) "  >  Starting server"
	$(LOG) "  >  $(PROJECTNAME) is available at $(ADDR)"
	@-$(GOBIN)/$(PROJECTNAME) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"

stop-server:
	$(LOG) "  >  Stopping server"
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true
	@-rm $(PID)

watch: ## watch: Run given command when code changes. e.g; make watch run="echo 'hey'"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) yolo -i . -e vendor -e bin -c "$(run)"

restart-server: stop-server start-server ## restart server

compile: ## compile: Compile the binary.
	@-touch $(STDERR)
	@-rm $(STDERR)
	@-$(MAKE) -s go-compile 2> $(STDERR)
	@cat $(STDERR) | sed -e '1s/.*/\nError:\n/'  | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

exec: ## exec: Run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
	$(LOG) "  >  Running exec"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(run)

clean: ## clean: Clean build files. Runs `go clean` internally.
	$(LOG) "  >  Cleaning build files"
	@-rm $(GOBIN)/$(PROJECTNAME) 2> /dev/null
	@-$(MAKE) go-clean

go-compile: go-get go-build ## GO compile
	$(LOG) "  >  GO compiling"

go-build: ## Build binary
	$(LOG) "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOBASE)/src/main/$(GOFILES)

go-generate: ## Generate dependency files
	$(LOG) "  >  Generating dependency files..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go generate $(generate)

go-get: ## Check for any missing dependencies
	$(LOG) "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(get)

go-install: ## GO install
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

go-clean: ## Clean build cache
	$(LOG) "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean
