// Code generated from Pkl module `kubecon.demo.infra.NodePoolConfig`. DO NOT EDIT.
package nodepool

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type NodePoolConfig struct {
	// nodepoolTemplates is a map keyed with a nodepool template name, and valued with the spec of desired NodePool.
	NodepoolTemplates map[string]*NodePoolSpec `pkl:"nodepoolTemplates"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a NodePoolConfig
func LoadFromPath(ctx context.Context, path string) (ret *NodePoolConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a NodePoolConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*NodePoolConfig, error) {
	var ret NodePoolConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
