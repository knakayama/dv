#!/usr/bin/env bash

set -xeou pipefail

if ! type "localstack" &>/dev/null; then
  # TODO: How can we pin a version with pipx?
  pipx install localstack
fi
