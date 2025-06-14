// Code generated from Pkl module `kubecon.demo.common.SelectorConfig`. DO NOT EDIT.
package operator

import (
	"encoding"
	"fmt"
)

type Operator string

const (
	In           Operator = "In"
	NotIn        Operator = "NotIn"
	Exists       Operator = "Exists"
	DoesNotExist Operator = "DoesNotExist"
)

// String returns the string representation of Operator
func (rcv Operator) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(Operator)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Operator.
func (rcv *Operator) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "In":
		*rcv = In
	case "NotIn":
		*rcv = NotIn
	case "Exists":
		*rcv = Exists
	case "DoesNotExist":
		*rcv = DoesNotExist
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid Operator`, str)
	}
	return nil
}
