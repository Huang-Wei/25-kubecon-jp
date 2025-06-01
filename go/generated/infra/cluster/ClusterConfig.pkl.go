// Code generated from Pkl module `kubecon.demo.infra.ClusterConfig`. DO NOT EDIT.
package cluster

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type ClusterConfig struct {
	// clusters define a list of desired Clusters.
	Clusters []*Cluster `pkl:"clusters"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a ClusterConfig
func LoadFromPath(ctx context.Context, path string) (ret *ClusterConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a ClusterConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*ClusterConfig, error) {
	var ret ClusterConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
