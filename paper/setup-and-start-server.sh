#!/bin/bash
sed -i "s/secret: PLACEHOLDER/secret: $SECRET/g" /data/config/paper-global.yml

exec /opt/minecraft/docker-entrypoint.sh