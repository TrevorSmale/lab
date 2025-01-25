# The Kube API Server

The **Kube API Server** is the central hub of communication in a Kubernetes cluster, acting as the **frontend** for the Kubernetes control plane. It facilitates communication with several key components of the cluster:

- **etcd cluster**
- **Kube-Scheduler**
- **Kubelet**
- **Controller Manager**

---

## D2 Diagram

I created **apiserver.d2** to illustrate this.

## etcd Cluster

The **Kube API Server** interacts with the **etcd cluster**, which serves as the **primary data store** for Kubernetes. It:
- Writes cluster state data to **etcd**.
- Reads cluster state data from **etcd** to determine the current state of the system.

---

## Kube-Scheduler

The **Kube API Server** enables two-way communication with the **Kube-Scheduler**:
- Sends **scheduling requests** when new workloads (Pods) need to be assigned to nodes.
- Receives **scheduling decisions** from the **Kube-Scheduler**, indicating where workloads should run.

---

## Kubelet

The **Kube API Server** manages communication with **Kubelet**, which runs on individual nodes. It:
- Sends **workload instructions** to deploy and manage containers on the nodes.
- Receives **Node and Pod status updates** from the **Kubelet**, providing insights into the health and status of workloads and nodes.

---

## Controller Manager

The **Kube API Server** coordinates with the **Controller Manager**, which manages various controllers responsible for maintaining the desired state of the cluster. It:
- Sends **controller instructions** to define the desired state (e.g., number of replicas for a Deployment).
- Receives **control loop updates** from the **Controller Manager**, which reflect changes made to align the cluster's actual state with the desired state.

