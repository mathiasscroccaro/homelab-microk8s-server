#!/bin/bash

tag_names=( 11 12 1subdomain )

kubernetes_registry="192.168.56.3:32000"

for tag_name in "${tag_names[@]}"
do
    TAG="$tag_name" envsubst < index_template.html > index.html
    
    sudo docker build -t nginx-$tag_name .
    sudo docker tag nginx-$tag_name $kubernetes_registry/nginx-$tag_name
    sudo docker push $kubernetes_registry/nginx-$tag_name
done
