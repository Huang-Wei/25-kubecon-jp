// Code generated from Pkl module `kubecon.demo.infra.AccountConfig`. DO NOT EDIT.
package account

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type AccountConfig struct {
	// accounts defines a list of Cloud accounts.
	Accounts []*Account `pkl:"accounts"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a AccountConfig
func LoadFromPath(ctx context.Context, path string) (ret *AccountConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a AccountConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*AccountConfig, error) {
	var ret AccountConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
