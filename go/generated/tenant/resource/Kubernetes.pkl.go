// Code generated from Pkl module `kubecon.demo.tenant.ResourceConfig`. DO NOT EDIT.
package resource

import "github.com/Huang-Wei/25-kubecon-jp/go/generated/tenant/selector"

type Kubernetes struct {
	// namespaces defines the Kubernetes namespaces to be created.
	Namespaces []string `pkl:"namespaces" json:"namespaces,omitempty" yaml:"namespaces,omitempty"`

	// selector defines the mapping Kubernetes clusters to create the resources.
	Selector []*selector.Requirment `pkl:"selector" json:"selector,omitempty" yaml:"selector,omitempty"`
}
