# kubectl-tap
A kubectl command to trigger immediate reevaluation of Kubernetes objects. It 
does this by updating an annotation--by default the key `tapped`--of the object.

This functionality generally is useful when developing or debugging 
controllers/operators, to avoid waiting for the periodic reconciliation loop 
(which can take minutes). By updating the annotation, the object will be 
enqueued and evaluated immediately by the corresponding controllers.  

## Installation
This project is a [Kubectl plugin](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/),
which means that it will be available as a subcommand in your local `kubectl` as
long as `kubectl-tap` is present under your PATH.

Install a prebuilt binary from one of the [releases](https://github.com/erwinvaneyk/kubectl-tap/releases):
```bash
VERSION=0.1.0
OS=darwin # options: [darwin, linux]
curl -fsSL -o kubectl-tap https://github.com/erwinvaneyk/kubectl-tap/releases/download/$VERSION/kubectl-tap-$OS-amd64
chmod u+x ./kubectl-tap
mv ./kubectl-tap /usr/local/bin/kubectl-tap
``` 

## Usage

Tap a single pod--for example the `kube-apiserver` pod in the `kube-system`:
```bash
$ kubectl tap -n kube-system pod/kube-apiserver
```

If you view the pod, an annotation should have been added to reflect the tap:
```bash
$ kubectl get -n kube-system pod/kube-apiserver -o yaml
apiVersion: v1
kind: Pod
metadata:
  annotations:
    tapped: "2020-01-14T11:06:02+01:00"
[...]
``` 

The command works similar to the regular commands of kubectl:
```bash
# Tap all deployments in kube-system
kubectl tap -n kube-system deployments --all 

# Tap all pods with the `component: kube-apiserver` label
kubectl tap -n kube-system -l component=kube-apiserver pod
```