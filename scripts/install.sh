#!/usr/bin/env bash

set -xeou pipefail

BIN_DIR="$(dirname ${0%/*})/bin"
[[ -d "$BIN_DIR" ]] || mkdir "$BIN_DIR"

if ! type "localstack" &>/dev/null; then
  # TODO: How can we pin a version with pipx?
  pipx install localstack
fi

if [[ ! -x "${BIN_DIR}/golangci-lint" ]]; then
  curl -sSfL "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" | sh -s "v1.49.0"
fi

if [[ ! -x "${BIN_DIR}/task" ]]; then
  sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b "$BIN_DIR" "v3.14.1"
fi

if [[ ! -x "${BIN_DIR}/gofumpt" ]]; then
  curl \
    -sSfL "https://github.com/mvdan/gofumpt/releases/download/v0.3.1/gofumpt_v0.3.1_darwin_amd64" \
    -o "${BIN_DIR}/gofumpt"
  chmod +x "${BIN_DIR}/gofumpt"
fi

ec_path="${BIN_DIR}/ec"
if [[ ! -x "$ec_path" ]]; then
  tmp_path="$(mktemp)"
  curl \
    -sSfL "https://github.com/editorconfig-checker/editorconfig-checker/releases/download/2.6.0/ec-darwin-amd64.tar.gz" \
    -o "$tmp_path"
  tar xzpvf "$tmp_path" -O > "$ec_path"
  chmod +x "$ec_path"
fi

gitleaks_path="${BIN_DIR}/gitleaks"
if [[ ! -x "$gitleaks_path" ]]; then
  tmp_path="$(mktemp)"
  curl \
    -sSfL "https://github.com/zricethezav/gitleaks/releases/download/v8.11.2/gitleaks_8.11.2_darwin_x64.tar.gz" \
    -o "$tmp_path"
  tar xzpvf "$tmp_path" "gitleaks"
  mv -v "gitleaks" "$gitleaks_path"
fi

goreleaser_path="${BIN_DIR}/goreleaser"
if [[ ! -x "$goreleaser_path" ]]; then
  tmp_path="$(mktemp)"
  curl \
    -sSfL "https://github.com/goreleaser/goreleaser/releases/download/v1.11.2/goreleaser_Darwin_x86_64.tar.gz" \
    -o "$tmp_path"
  tar xzpvf "$tmp_path" "goreleaser"
  mv -v "goreleaser" "$goreleaser_path"
fi
