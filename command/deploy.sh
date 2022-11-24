#!/bin/env bash

docker stop esignatures
docker rm esignatures
docker rmi esignatures
git pull
docker build -t esignatures .
docker container create --name esignatures -e PORT=443 -e INSTANCE_ID="smartsign signatures" -p 443:443 esignatures
docker container start esignatures

