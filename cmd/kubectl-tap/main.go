package main

import (
	"github.com/erwinvaneyk/kubectl-tap/pkg/cmd"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-tap", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCmdTap(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}