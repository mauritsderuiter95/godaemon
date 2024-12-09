#!/bin/sh

rm -rf /config/godaemon/cmd
rm -rf /config/godaemon/pkg
cp /app/src/* /config/godaemon

cd /config/godaemon

set -e

while true; do
  go build -o ./godaemon-app ./cmd/godaemon/.
  ./godaemon-app
done