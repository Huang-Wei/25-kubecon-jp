// Code generated from Pkl module `kubecon.demo.infra.AccountConfig`. DO NOT EDIT.
package account

type Account struct {
	// account_id specifies the account id that manages the Cluster(s). It varies upon 'provider':
	// - aws: a 12-digit number.
	// - gcp: a random string that's also known as "project name".
	AccountID string `pkl:"accountID"`

	// cloudProvider specfies the cloud provider name.
	CloudProvider string `pkl:"cloudProvider"`

	// tags is a map used to match Resource's selector.
	Tags map[string]string `pkl:"tags"`
}
