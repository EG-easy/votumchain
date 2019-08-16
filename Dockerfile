# Simple usage with a mounted data directory:
# > docker build -t votum .
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.votumd:/root/.votumd -v ~/.votumcli:/root/.votumcli votum votumd init
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.votumd:/root/.votumd -v ~/.votumcli:/root/.votumcli votum votumd start
FROM golang:alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python

# Set working directory for the build
WORKDIR /go/src/github.com/EG-easy/votumchain

# Add source files
COPY . .

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES && \
    export GO111MODULE=on && \
    make install

# # Final image
# FROM alpine:edge
#
# # Install ca-certificates
# RUN apk add --update ca-certificates
# WORKDIR /root
#
# # Copy over binaries from the build-env
# COPY --from=build-env /go/bin/votumd /usr/bin/votumd
# COPY --from=build-env /go/bin/votumcli /usr/bin/votumcli
#
# # Run votumd by default, omit entrypoint to ease using container with votumcli
# # CMD ["votumd start"]
