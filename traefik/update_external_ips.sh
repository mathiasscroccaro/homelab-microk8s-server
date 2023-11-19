INTERFACE="wlp2s0"
IPV4=$(ip addr show dev $INTERFACE scope global | sed -e's/^.*inet6 \([^ ]*\)\/.*$/\1/;t;d')
IPV6=$(ip -6 addr show dev $INTERFACE scope global | sed -e's/^.*inet6 \([^ ]*\)\/.*$/\1/;t;d')
microk8s kubectl patch svc traefik -n traefik --type merge -p '{"spec":{"externalIPs":["'"${IPV4}"'","'"${IPV6}"'"]}}'