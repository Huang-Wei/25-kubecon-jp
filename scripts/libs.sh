#!/usr/bin/env bash

function must_on_path() {
  local program="$1"
  if ! command -v $program &> /dev/null; then
    echo "$program is not installed, abort."
    exit 1
  fi
}

function pkl_go_codegen() {
  if ! command -v pkl-gen-go &> /dev/null; then
    echo "Install pkl-gen-go to run this script: go install github.com/apple/pkl-go/cmd/pkl-gen-go@latest"
    exit 1
  fi

  rm -rf ${SCRIPT_ROOT}/go/generated/*

  find ${SCRIPT_ROOT}/schema -type f  -name "*.pkl" | sort | while read file; do
    echo "===$file==="
    pkl-gen-go $file
  done
}
