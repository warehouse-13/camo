#!/bin/bash

PASS="${1:-dev}"

docker run -e CLEAR_PASSWORD="$PASS" ghcr.io/warehouse-13/camo:latest
