#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname $(dirname "${BASH_SOURCE[@]}"))
source "${SCRIPT_ROOT}/scripts/libs.sh"

# Verify all tenants' input
find ${SCRIPT_ROOT}/tenants -mindepth 1 -maxdepth 1 -type d | sort | while read tenant; do
  echo "🔍 Validating tenant '$tenant'"
  find $tenant -type f -name "*.pkl" | sort | while read pkl_file; do
    echo -n "   - Validating $pkl_file"
    pkl eval $pkl_file > /dev/null
    [[ $? -eq 0 ]] && echo " ✅" || echo " ❌"
  done
done
