#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <docker_container>"
  exit 1
fi

DOCKER_CONTAINER=$1

docker cp uploads/ $DOCKER_CONTAINER:/app

docker cp passport_scanner.db $DOCKER_CONTAINER:/app
