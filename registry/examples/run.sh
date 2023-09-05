#!/bin/bash

kubernetes_registry_url="192.168.15.71:32000"

sudo docker build -t nginx-test-registry .
sudo docker tag nginx-test-registry $kubernetes_registry_url/nginx-test-registry
sudo docker push $kubernetes_registry_url/nginx-test-registry
