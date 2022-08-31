#!/usr/bin/env bash -xeou pipefail

ROOT_DIR="$(cd ${0%/*}/.. && pwd -P)"

if [[ ! -x "${ROOT_DIR}/bin/golangci-lint" ]]; then
  curl -sSfL "https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh" | sh -s "v1.49.0"
fi

if [[ ! -x "${ROOT_DIR}/bin/task" ]]; then
  sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b "${ROOT_DIR}/bin" "v3.14.1"
fi

if [[ ! -x "${ROOT_DIR}/bin/gofumpt" ]]; then
  curl \
    -sSfL "https://github.com/mvdan/gofumpt/releases/download/v0.3.1/gofumpt_v0.3.1_darwin_amd64" \
    -o "${ROOT_DIR}/bin/gofumpt"
  chmod +x "${ROOT_DIR}/bin/gofumpt"
fi

ec_path="${ROOT_DIR}/bin/ec"
if [[ ! -x "$ec_path" ]]; then
  tmp_path="$(mktemp)"
  curl \
    -sSfL "https://github.com/editorconfig-checker/editorconfig-checker/releases/download/2.6.0/ec-darwin-amd64.tar.gz" \
    -o "$tmp_path"
  tar xzpvf "$tmp_path" -O > "$ec_path"
  chmod +x "$ec_path"
fi
