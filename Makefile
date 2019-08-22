#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

export GO111MODULE = on

# process build tags

build_tags = netgo
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=votum \
		  -X github.com/cosmos/cosmos-sdk/version.ServerName=votumd \
		  -X github.com/cosmos/cosmos-sdk/version.ClientName=votumcli \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

#### Command List ####

all: lint install

install: go.sum
		go install $(BUILD_FLAGS) ./cmd/votumd
		go install $(BUILD_FLAGS) ./cmd/votumcli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

build:
	go build $(BUILD_FLAGS) -o build/votumd ./cmd/votumd
	go build $(BUILD_FLAGS) -o build/votumcli ./cmd/votumcli

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
