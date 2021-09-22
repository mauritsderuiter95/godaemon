#!/bin/sh

set -e

go build -o ./godaemon-app ./cmd/godaemon/.

./godaemon-app