// Code generated from Pkl module `kubecon.demo.tenant.ResourceConfig`. DO NOT EDIT.
package resource

import "github.com/Huang-Wei/25-kubecon-jp/go/generated/common/selector"

type Bucket struct {
	Name string `pkl:"name"`

	Region string `pkl:"region"`

	Ttl *string `pkl:"ttl" json:"ttl,omitempty" yaml:"ttl,omitempty"`

	Selector []*selector.Requirment `pkl:"selector" json:"selector,omitempty" yaml:"selector,omitempty"`
}
