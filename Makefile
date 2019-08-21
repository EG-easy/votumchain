#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')

export GO111MODULE = on

# process build tags

build_tags = netgo
UNAME_S = $(shell uname -s)
ifeq ($(UNAME_S),OpenBSD)
	$(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
else
	GCC = $(shell command -v gcc 2> /dev/null)
	ifeq ($(GCC),)
		$(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
	else
		build_tags += ledger
	endif
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=gaia \
		  -X github.com/cosmos/cosmos-sdk/version.ServerName=gaiad \
		  -X github.com/cosmos/cosmos-sdk/version.ClientName=gaiacli \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

#### Command List ####

all: lint install

install: go.sum
		go install ./cmd/votumd
		go install ./cmd/votumcli
		@echo VERSION

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

build:
	go build -o build/votumd ./cmd/votumd
	go build -o build/votumcli ./cmd/votumcli

#docker用のbinaryを作る
build-linux: go.sum
	GOOS=linux GOARCH=amd64 $(MAKE) build

# Makefile for the "votumdnode" docker image.
build-docker:
	$(MAKE) -C docker/local

# Run a 4-node testnet locally
localnet-start: build-linux localnet-stop
	@if ! [ -f build/node0/votumd/config/genesis.json ]; then docker run --rm -v $(CURDIR)/build:/votumd:Z votumchain/votumdnode testnet --v 4 -o . --starting-ip-address 192.168.10.2 ; fi
	docker-compose up -d

# Stop testnet
localnet-stop:
	docker-compose down
