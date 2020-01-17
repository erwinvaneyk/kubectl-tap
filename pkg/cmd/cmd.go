package cmd

import (
	"fmt"
	"github.com/erwinvaneyk/kubectl-tap/pkg/version"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/dynamic"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"strings"
	"time"

	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	tapExample = `
	# tap deployment foo to trigger reevaluation of the deployment controller.
	%[1]s deployment/foo
`
	defaultTapKey = "tapped"
)

type TapOptions struct {
	genericclioptions.IOStreams
	configFlags *genericclioptions.ConfigFlags

	args               []string
	tapKey             string
	targetResourceType string
	targetResourceName string
	targetNamespace    string
	labelSelector      string
	allResources       bool
}

// NewTapOptions provides an instance of TapOptions with default values
func NewTapOptions(streams genericclioptions.IOStreams) *TapOptions {
	return &TapOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
		IOStreams:   streams,
		tapKey:      defaultTapKey,
	}
}

// NewCmdTap provides a cobra command wrapping TapOptions
func NewCmdTap(streams genericclioptions.IOStreams, version version.Info) *cobra.Command {
	o := NewTapOptions(streams)

	cmd := &cobra.Command{
		Use:          "tap TYPE[.VERSION][.GROUP]/NAME [flags]",
		Short:        "Trigger watching controller to reevaluate the resource(s).",
		Example:      fmt.Sprintf(tapExample, "kubectl"),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
		Version: version.String(),
	}

	cmd.Flags().StringVar(&o.tapKey, "key", o.tapKey, "The key of annotation to update for the resource.")
	cmd.Flags().StringVarP(&o.labelSelector, "selector", "l", o.labelSelector, "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	cmd.Flags().BoolVar(&o.allResources, "all", o.allResources, "Tap all resources in the namespace of the specified resource types.")

	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

// Complete sets all information required for updating the current context
func (o *TapOptions) Complete(cmd *cobra.Command, args []string) error {
	var err error

	o.args = args

	o.targetNamespace, _, err = o.configFlags.ToRawKubeConfigLoader().Namespace()
	if err != nil {
		return err
	}

	return nil
}

// Validate ensures that all required arguments and flag values are provided
func (o *TapOptions) Validate() error {
	return nil
}

func (o *TapOptions) Run() error {
	restCfg, err := o.configFlags.ToRESTConfig()
	if err != nil {
		return err
	}

	dynamicClient, err := dynamic.NewForConfig(restCfg)
	if err != nil {
		return err
	}

	resp := cmdutil.NewFactory(o.configFlags).NewBuilder().
		Unstructured().
		NamespaceParam(o.targetNamespace).
		DefaultNamespace().
		ResourceTypeOrNameArgs(o.allResources, o.args...).
		LabelSelector(o.labelSelector).
		SingleResourceType().
		ContinueOnError().
		Latest().
		Flatten().
		Do()

	return resp.Visit(func(info *resource.Info, err error) error {
		unstructuredObj := &unstructured.Unstructured{}
		unstructuredObj.Object, err = runtime.DefaultUnstructuredConverter.ToUnstructured(info.Object)
		if err != nil {
			return err
		}

		updatedObj := o.tap(unstructuredObj)

		// Update the tapped object
		updateResp, err := dynamicClient.Resource(info.ResourceMapping().Resource).
			Namespace(o.targetNamespace).
			Update(updatedObj, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
		fmt.Printf("%s/%s tapped\n", strings.ToLower(updateResp.GetKind()), updateResp.GetName())

		return nil
	})
}

func (o *TapOptions) tap(obj *unstructured.Unstructured) *unstructured.Unstructured {
	updatedObj := obj.DeepCopy()
	annotations := updatedObj.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}
	annotations[o.tapKey] = time.Now().Format(time.RFC3339)
	updatedObj.SetAnnotations(annotations)
	return updatedObj
}
