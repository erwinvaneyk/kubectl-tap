module github.com/erwinvaneyk/kubectl-tap

go 1.13

require (
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	k8s.io/apimachinery v0.17.0
	k8s.io/cli-runtime v0.17.0
	k8s.io/client-go v0.17.0
	k8s.io/kubectl v0.17.0
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
	k8s.io/api => k8s.io/api v0.0.0-20200113233642-3946df5ca773
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200113233504-44bd77c24ef9
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20200113235527-4c0b167ce833
	k8s.io/client-go => k8s.io/client-go v0.0.0-20200113233857-bcaa73156d59
)
