#!/usr/bin/env sh

##
## Input parameters
##
BINARY=/votumd/${BINARY:-votumd}
ID=${ID:-0}
LOG=${LOG:-votumd.log}

##
## Assert linux binary
##
if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'votumd' E.g.: -e BINARY=votumd_my_test_version"
	exit 1
fi
BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"
if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi
##
## Run binary with all parameters
##
export VOTUMDHOME="/votumd/node${ID}/votumd"

if [ -d "`dirname ${VOTUMDHOME}/${LOG}`" ]; then
  "$BINARY" --home "$VOTUMDHOME" "$@" | tee "${VOTUMDHOME}/${LOG}"
else
  "$BINARY" --home "$VOTUMDHOME" "$@"
fi

chmod 777 -R /votumd

