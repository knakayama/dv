#!/usr/bin/env bash

set -xeuo pipefail

task lint

git add $(git diff --cached --name-only --diff-filter=ACMR)
