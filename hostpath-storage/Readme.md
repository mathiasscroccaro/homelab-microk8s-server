# How to create a External HDD host-storage

1. Automatically mount your external hard drive
2. Enable the hostpath-storage `microk8s enable hostpath-storage`
2. Configure your storage class 

## Automatically mount your external hard drive

1. Create folder where the external hard drive will be mounted

```
sudo mkdir /mnt/external-hdd
```

2. Get the UUID of the partition

```
# sudo blkid
...
/dev/sda4: LABEL="writable" UUID="fa33e88e-a697-4cef-95f5-8eb90dbc3394" BLOCK_SIZE="4096" TYPE="ext4" PARTUUID="9f10f025-addd-af4e-8ac1-67400bc1a1a0"
/dev/sda2: SEC_TYPE="msdos" LABEL_FATBOOT="ESP" LABEL="ESP" UUID="F7DB-4D56" BLOCK_SIZE="512" TYPE="vfat" PARTLABEL="Appended2" PARTUUID="c3b09d5c-3f01-457b-ac5c-206e8818bfa0"
/dev/sda3: PARTLABEL="Gap1" PARTUUID="c3b09d5c-3f01-457b-ac5d-206e8818bfa0"
/dev/sda1: BLOCK_SIZE="2048" UUID="2023-02-17-21-57-15-00" LABEL="Ubuntu-Server 22.04.2 LTS amd64" TYPE="iso9660" 
...
```

3. Install the UUID in the `/etc/fstab` file

```
UUID=fa33e88e-a697-4cef-95f5-8eb90dbc3394	/mnt/external-hdd	ext4	defaults	0	2
```

## Configure your storage class 

1. Create a storage file and apply it

```
kubectl apply -f external-hdd-sc.yaml
```

2. To test it, run it

```
microk8s kubectl apply -f - <<EOF
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: test-pvc-hdd
spec:
  storageClassName: external-hdd-hostpath
  accessModes: [ReadWriteOnce]
  resources: { requests: { storage: 1Gi } }
---
apiVersion: v1
kind: Pod
metadata:
  name: test-nginx-hdd
spec:
  volumes:
    - name: pvc
      persistentVolumeClaim:
        claimName: test-pvc-hdd
  containers:
    - name: nginx
      image: nginx
      ports:
        - containerPort: 80
      volumeMounts:
        - name: pvc
          mountPath: /usr/share/nginx/html
EOF
```

