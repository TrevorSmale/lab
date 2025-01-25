# Anatomy of a Config File
## 3 Parts

**1. Metadata**
**2. Specification**
**3. Status (Generated & Adhered)**

Best practice is to store with code or...
store them seperately in own repo

```yaml
nginx-deployment.yaml
apiVersion:
kind:
metadate
    name:
    labels:
spec:
    replicas:
    selector:
    template:
status: updated continuously from etcd
```

etcd holds the status of any Kubernetes deployment

Docker & Containerd

Though I do not really use Docker and prefer Podman. 