// Code generated from Pkl module `kubecon.demo.tenant.SelectorConfig`. DO NOT EDIT.
package selector

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type SelectorConfig struct {
	Selector []*Requirment `pkl:"selector"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a SelectorConfig
func LoadFromPath(ctx context.Context, path string) (ret *SelectorConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a SelectorConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*SelectorConfig, error) {
	var ret SelectorConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
