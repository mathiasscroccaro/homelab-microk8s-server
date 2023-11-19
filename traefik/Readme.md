## How to install Traefik

```
microk8s enable community
microk8s enable traefik
```

```
# Install Traefik Resource Definitions:
kubectl apply -f https://raw.githubusercontent.com/traefik/traefik/v2.10/docs/content/reference/dynamic-configuration/kubernetes-crd-definition-v1.yml

# Install RBAC for Traefik:
kubectl apply -f https://raw.githubusercontent.com/traefik/traefik/v2.10/docs/content/reference/dynamic-configuration/kubernetes-crd-rbac.yml
```

# Change the load balancer port of the traefik

Use the command `kubectl edit svc/traefik -n traefik` to change the web port to 8000 and websecure to 8443.

Also, with the same command, add to the ipFamilies list the value "ipv6" and change the ipFamilyPolicy to "RequireDualStack".

Copy the script `update_external_ips.sh` to the cluster and make it run every 30 minutes `*/30 * * * * /opt/update_external_ips/update_external_ips.sh`

```
INTERFACE="wlp2s0"
IPV4=$(ip addr show dev $INTERFACE scope global | sed -e's/^.*inet6 \([^ ]*\)\/.*$/\1/;t;d')
IPV6=$(ip -6 addr show dev $INTERFACE scope global | sed -e's/^.*inet6 \([^ ]*\)\/.*$/\1/;t;d')
microk8s kubectl patch svc traefik -n traefik --type merge -p '{"spec":{"externalIPs":["192.168.15.71","2804:1b3:a7c2:315c:3aca:73ff:fe0e:8fd2"]}}'
```