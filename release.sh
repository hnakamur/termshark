#!/bin/sh
if [ $# -ne 1 ]; then
  echo 'Usage: ./release.sh TERMSHARK_VERSION' >&2 
  echo 'Example: ./release.sh 1.0.0' >&2 
  exit 2
fi
export DOCKER_BUILDKIT=1
docker build -t termshark --build-arg TERMSHARK_VERSION=$1 --secret id=github_token,src=.github_token .
