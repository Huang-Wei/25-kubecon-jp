package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Huang-Wei/25-kubecon-jp/go/generated/tenant/resource"
)

func main() {
	ret, err := resource.LoadFromPath(context.Background(), "/Users/weih/projects/25-kubecon-jp/tenants/foo/dev/resource.pkl")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, ns := range ret.Kubernetes.Namespaces {
		fmt.Println(ns)
	}
}
