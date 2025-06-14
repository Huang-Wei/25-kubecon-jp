// Code generated from Pkl module `kubecon.demo.common.SelectorConfig`. DO NOT EDIT.
package selector

import (
	"github.com/Huang-Wei/25-kubecon-jp/go/generated/common/selector/key"
	"github.com/Huang-Wei/25-kubecon-jp/go/generated/common/selector/operator"
)

type Requirment struct {
	// key is the label key that the selector applies to.
	Key key.Key `pkl:"key"`

	// operator represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator operator.Operator `pkl:"operator"`

	// values is an array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty.
	Values []string `pkl:"values"`
}
