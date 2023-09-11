# How to remove all unused docker images from your k8s cluster's cache

1. [Install crictl](https://github.com/kubernetes-sigs/cri-tools/blob/edf14e37007994d69f9b8cb3b1483a79b365b8eb/docs/crictl.md#install-crictl)

2. Configure the crictl to use the docker sock

    + Create or edit the file `/etc/crictl.yaml`
    + Add the following content to it:

```
runtime-endpoint: unix:///var/snap/microk8s/common/run/containerd.sock
image-endpoint: unix:///var/snap/microk8s/common/run/containerd.sock
timeout: 10
debug: false
```

3. Check if the crictl is working: `crictl images`

4. Remove unused images `crictl rmi --prune`

5. Add the previous command to the root crontab:

```
* */1 * * * crictl rmi --prune 
```