SHELL += -eu

BLUE  := \033[0;34m
GREEN := \033[0;32m
YELLOW := \033[0;33m
RED   := \033[0;31m
NC    := \033[0m

GO111MODULE := on
GOPATH ?= ${HOME}/.gvm/gos/go1.14
GO_BIN := ${GOPATH}/bin
GOPRIVATE := github.com/kevinmichaelchen

.PHONY: all
all:
	docker-compose up -d
	echo Sleeping for a bit
	sleep 7
	env \
	  GO111MODULE=${GO111MODULE} \
	  go run main.go

include makefiles/*.mk
