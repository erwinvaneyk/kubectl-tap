# kubectl-tap
This project is an example of a simple kubectl plugin. It showcases typical 
components (cli-runtime, cobra, apimachinery, ...) and tooling (goreleaser, krew) 
when working within the Kubernetes ecosystem.

The `tap` command is roughly equivalent to `kubectl annotate pod tapped=$(date)` 
with the use case to trigger a new controller evaluation of the target objects. 
It does this by updating an annotation--by default the key `tapped`--of the 
object.

## Installation
The command is a [Kubectl plugin](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/),
which means that it will be available as a subcommand in your local `kubectl` as
long as `kubectl-tap` is present under your PATH.

Install a prebuilt binary from one of the [releases](https://github.com/erwinvaneyk/kubectl-tap/releases):
```bash
OS=darwin # options: [darwin, linux, windows]
curl -fsSL -o kubectl-tap https://github.com/erwinvaneyk/kubectl-tap/releases/download/v0.1.0/kubectl-tap-$OS-amd64
chmod u+x ./kubectl-tap
mv ./kubectl-tap /usr/local/bin
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
