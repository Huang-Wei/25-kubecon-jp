package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Huang-Wei/25-kubecon-jp/go/generated/tenant/resource"
)

func main() {
	pkls := []string{"tenants/bar/prod/resource.pkl", "tenants/foo/dev/resource.pkl"}

	for _, pkl := range pkls {
		fmt.Printf("===%s===\n", pkl)
		ret, err := resource.LoadFromPath(context.Background(), pkl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, ns := range ret.Kubernetes.Namespaces {
			fmt.Println(ns)
		}
	}
}
