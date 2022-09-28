#!/usr/bin/env bash

set -xeuo pipefail

task lint go:vuln

git add $(git diff --cached --name-only --diff-filter=ACMR)
