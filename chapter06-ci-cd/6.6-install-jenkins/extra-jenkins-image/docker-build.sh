#!/bin/bash

version=2.425

docker buildx build --platform linux/amd64 -t 3dsinteractive/jenkins:$version .
docker push 3dsinteractive/jenkins:$version