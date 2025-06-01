// Code generated from Pkl module `kubecon.demo.tenant.ResourceConfig`. DO NOT EDIT.
package resource

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type ResourceConfig struct {
	Kubernetes *Kubernetes `pkl:"kubernetes" json:"kubernetes,omitempty" yaml:"kubernetes,omitempty"`

	Buckets []*Bucket `pkl:"buckets" json:"buckets,omitempty" yaml:"buckets,omitempty"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a ResourceConfig
func LoadFromPath(ctx context.Context, path string) (ret *ResourceConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a ResourceConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*ResourceConfig, error) {
	var ret ResourceConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
