package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func main() {
	res := "posts"
	configFlag := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()
	matchVersionFlags := cmdutil.NewMatchVersionFlags(configFlag)
	m, err := cmdutil.NewFactory(matchVersionFlags).ToRESTMapper()
	if err != nil {
		fmt.Printf("gettint rest mapper %s", err.Error())
	}
	gvr, err := m.ResourceFor(schema.GroupVersionResource{
		Resource: res,
	})

	if err != nil {
		fmt.Printf("getting GVR for res %s", err.Error())
	}
	fmt.Printf("compplete GVR is Group: %s, version %s, resource %s", gvr.Group, gvr.Version, gvr.Resource)

}
