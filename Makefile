all: lint install

install: go.sum
		go install ./cmd/votumd
		go install ./cmd/votumcli

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
