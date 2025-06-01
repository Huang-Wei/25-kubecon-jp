// Code generated from Pkl module `kubecon.demo.infra.NodePoolConfig`. DO NOT EDIT.
package nodepool

type NodePoolSpec struct {
	// tainted specified whether the Nodes of this NodePool will be tainted with
	// "<name>=true:NoSchedule". Defaults to true.
	Tainted bool `pkl:"tainted"`

	// min_size defines the minimum size of the NodePool.
	MinSize *int `pkl:"minSize" json:"minSize,omitempty" yaml:"minSize,omitempty"`

	// max_size defines the minimum size of the NodePool.
	MaxSize *int `pkl:"maxSize" json:"maxSize,omitempty" yaml:"maxSize,omitempty"`

	// instanceTypes specfies the types of instances that can be used in this Nodepool.
	InstanceTypes []string `pkl:"instanceTypes"`

	// tags defines a seires of key:value pairs for the NodePool.
	Tags map[string]string `pkl:"tags" json:"tags,omitempty" yaml:"tags,omitempty"`
}
