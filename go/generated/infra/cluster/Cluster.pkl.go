// Code generated from Pkl module `kubecon.demo.infra.ClusterConfig`. DO NOT EDIT.
package cluster

import "github.com/Huang-Wei/25-kubecon-jp/go/generated/infra/nodepool"

// Cluster describe the spec of a desired Cluster.
type Cluster struct {
	Name string `pkl:"name"`

	// k8sVersion describes the desired Kubernetes version of the cluster(s).
	K8sVersion string `pkl:"k8sVersion"`

	// nodePools defines a list of NodePools.
	// Each NodePool derives the spec from pre-configured node templates, and optinally
	// customize the spec.
	NodePools []*nodepool.NodePoolSpec `pkl:"nodePools"`
}
