//go:build tools

package ports

//nolint
import (
	_ "github.com/golang/mock/mockgen/model" // github.com/golang/mock/mockgen@latest
	_ "golang.org/x/tools/imports"           // golang.org/x/tools/cmd/goimports@latest
)
