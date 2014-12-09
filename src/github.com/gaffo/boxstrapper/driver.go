package boxstrapper

import (
	// "os"
)

type Driver interface {
	AddPackage(packageName string) error
}

type Storage interface {
	ReadPackages() (string, error)
	WritePackages(contents, reason string) error
}