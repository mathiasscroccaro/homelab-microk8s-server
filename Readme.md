# Private microk8s cluster

Here you'll find the configurations I used to build my private local k8s server.

Each folder is a different service

## After install

After install the microk8s, you need to copy the file `/var/snap/microk8s/current/credentials/client.config` to your client folder `~/.kube/`. You'll also need to rename it to `microk8s.config`.

## How to configure multiple kubernetes context

```
echo "export KUBECONFIG=/home/mathias/.kube/k3s.config:/home/mathias/.kube/microk8s.config" >> ~/.profile
source ~/.profile
```

## How to set context

Search the cluster name:

```
kubectl config get-contexts
```

and set it:

```
kubectl config use-context <context-name>
```

## Steps to replicate the microk8s cluster

1. Configure the router to port fowarding the external port 80 to your local machine port
2. Install the dynamic DNS software 
3. Install the microk8s in your server
4. Enable the hostpath StorageService `microk8s enable hostpath-storage`
5. Install the external-hdd storage class, using the content of the folder `./hostpath-storage`
6. Install the rest of the services
