// Code generated from Pkl module `kubecon.demo.tenant.SelectorConfig`. DO NOT EDIT.
package key

import (
	"encoding"
	"fmt"
)

type Key string

const (
	Env           Key = "env"
	Geo           Key = "geo"
	ClusterType   Key = "cluster_type"
	CloudProvider Key = "cloud_provider"
)

// String returns the string representation of Key
func (rcv Key) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(Key)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Key.
func (rcv *Key) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "env":
		*rcv = Env
	case "geo":
		*rcv = Geo
	case "cluster_type":
		*rcv = ClusterType
	case "cloud_provider":
		*rcv = CloudProvider
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid Key`, str)
	}
	return nil
}
