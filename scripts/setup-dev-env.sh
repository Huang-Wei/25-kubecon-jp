#!/usr/bin/env bash

go install github.com/apple/pkl-go@latest
go install k8s.io/code-generator/cmd/deepcopy-gen@v0.32.4
brew install pkl yq
