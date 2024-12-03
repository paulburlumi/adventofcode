//go:build tools

// Adding tools as dependencies. See https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
package tools

import (
	_ "github.com/golang/mock/mockgen"
	_ "golang.org/x/tools/cmd/godoc"
	_ "golang.org/x/tools/cmd/stringer"
	_ "golang.org/x/vuln/cmd/govulncheck"
	_ "mvdan.cc/gofumpt"
)