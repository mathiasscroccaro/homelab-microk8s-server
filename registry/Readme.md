# How to configure a private registry

Enable the private registry:

```
microk8s enable registry
```

## How to use external hard drive as PVC for the registry

1. Deploy the external-hdd-sc, presented in the directory `../hostpath-storage`:

```
kubectl apply -f ../hostpath-storage/external-hdd-sc.yaml
```

2. Create the PVC to use the external hard drive:

```
kubectl apply -f pvc-registry.yaml
```

3. Patch the deployed registry to use the just created PVC:

```
kubectl patch deployment.apps/registry -n container-registry -p '{"spec":{"template":{"spec":{"volumes":[{"name":"registry-data","persistentVolumeClaim":{"claimName":"registry-pvc-hdd"}}]}}}}'
```

## How to test the deployed private registry

