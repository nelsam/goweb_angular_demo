package settings

import (
	"os"
	"path"
)

const (
	GoPathEnv = "GOPATH"
)

var (

	// ProjectImportPath is the import path to the base package for
	// this project.
	ProjectImportPath = "github.com/nelsam/goweb_angular_demo"

	// ProjectPath is the path to this project.
	ProjectPath = path.Join(os.Getenv(GoPathEnv), "src", ProjectImportPath)
)
