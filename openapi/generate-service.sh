#!/bin/sh
APIDIR=$(dirname "$0")
cd ${APIDIR}
${APIDIR}/../node_modules/.bin/openapi-generator-cli generate -c generate-config.yaml
