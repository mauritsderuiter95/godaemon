#!/bin/sh

rm -rf /config/godaemon/cmd
rm -rf /config/godaemon/pkg
mv /data/app/src/* /config/godaemon

cd /config/godaemon

while true; do
  go build -o ./godaemon-app ./cmd/godaemon/.
  ./godaemon-app
done