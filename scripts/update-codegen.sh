#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname $(dirname "${BASH_SOURCE[@]}"))
source "${SCRIPT_ROOT}/scripts/libs.sh"

must_on_path pkl-gen-go
pkl_go_codegen

# echo

# must_on_path deepcopy-gen

# k8s_codegen "${SCRIPT_ROOT}/go/generated/infra"
# k8s_codegen "${SCRIPT_ROOT}/go/generated/tenant"
