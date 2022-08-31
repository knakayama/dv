#!/usr/bin/env bash -xeou pipefail

TOOLS_DIR="$(cd ${0%/*}/../tools && pwd -P)"

if ! type "localstack" &>/dev/null; then
  # TODO: How can we pin a version with pipx?
  pipx install localstack
fi

if [[ ! -x "${TOOLS_DIR}/golangci-lint" ]]; then
  curl -sSfL "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" | sh -s "v1.49.0"
fi

if [[ ! -x "${TOOLS_DIR}/task" ]]; then
  sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b "$TOOLS_DIR" "v3.14.1"
fi

if [[ ! -x "${TOOLS_DIR}/gofumpt" ]]; then
  curl \
    -sSfL "https://github.com/mvdan/gofumpt/releases/download/v0.3.1/gofumpt_v0.3.1_darwin_amd64" \
    -o "${TOOLS_DIR}/gofumpt"
  chmod +x "${TOOLS_DIR}/gofumpt"
fi

ec_path="${TOOLS_DIR}/ec"
if [[ ! -x "$ec_path" ]]; then
  tmp_path="$(mktemp)"
  curl \
    -sSfL "https://github.com/editorconfig-checker/editorconfig-checker/releases/download/2.6.0/ec-darwin-amd64.tar.gz" \
    -o "$tmp_path"
  tar xzpvf "$tmp_path" -O > "$ec_path"
  chmod +x "$ec_path"
fi

gitleaks_path="${TOOLS_DIR}/gitleaks"
if [[ ! -x "$gitleaks_path" ]]; then
  tmp_path="$(mktemp)"
  curl \
    -sSfL "https://github.com/zricethezav/gitleaks/releases/download/v8.11.2/gitleaks_8.11.2_darwin_x64.tar.gz" \
    -o "$tmp_path"
  tar xzpvf "$tmp_path" "gitleaks"
  mv -v "gitleaks" "$gitleaks_path"
fi
