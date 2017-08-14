#!/bin/bash
set -e
set -x

BASE_URL=http://localhost:9000
mkdir -p .test

# test GET
curl --silent $BASE_URL/hello -o ./.test/hello
test $(jq -r '.msg' < .test/hello) = hello

curl --silent $BASE_URL/issues -o ./.test/issues
test $(jq -r '. | length' < .test/issues) = 2
test "$(jq -r '.[1].Name' < .test/issues)" = "I need another help"

# test POST
curl --silent -X POST $BASE_URL/hello -o ./.test/post_hello
test "$(jq -r '.msg' < .test/post_hello)" = "post hello"

echo "ok"

