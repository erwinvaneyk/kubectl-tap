package main

import (
	"github.com/erwinvaneyk/kubectl-tap/pkg/cmd"
	versionpkg "github.com/erwinvaneyk/kubectl-tap/pkg/version"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
)

// Variables filled with ldflags:
var (
	version string
	commit  string
	date    string
)

func main() {
	flags := pflag.NewFlagSet("kubectl-tap", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCmdTap(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}, versionpkg.Info{
		Version:   version,
		Commit:    commit,
		BuildDate: date,
	})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
