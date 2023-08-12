# Private microk8s and k3s services

## After install

After install the k3s, you need to copy the file `/etc/rancher/k3s/k3s.yaml` to your client folder `~/.kube/`. You'll also need to rename it to `k3s.config`.

After install the microk8s, you need to copy the file `/var/snap/microk8s/current/credentials/client.config` to your client folder `~/.kube/`. You'll also need to rename it to `microk8s.config`.

## How to configure multiple kubernetes context

```
echo "export KUBECONFIG=/home/mathias/.kube/k3s.config:/home/mathias/.kube/microk8s.config" >> ~/.profile
source ~/.profile
```

## How to set context

Search the cluster name:

`kubectl config get-contexts`

and set it:

`kubectl config use-context <context-name>`
