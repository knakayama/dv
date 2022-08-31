#!/usr/bin/env bash -xeuo pipefail

./bin/task lint

git add $(git diff --cached --name-only --diff-filter=ACMR)
