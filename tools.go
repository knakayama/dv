//go:build tools

package tools

import (
	_ "github.com/editorconfig-checker/editorconfig-checker/cmd/editorconfig-checker"
	_ "github.com/go-task/task/v3/cmd/task"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/zricethezav/gitleaks/v8"
	_ "golang.org/x/vuln/cmd/govulncheck"
	_ "mvdan.cc/gofumpt"
)
