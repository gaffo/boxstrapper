package boxstrapper

import (
// "os"
)

type Driver interface {
	AddPackage(packageName string) error
}

type Storage interface {
	ReadOpsfile() (string, error)
	WriteOpsfile(contents string, reason string) error
}

type OperationsStorage interface {
	ReadOperations() ([]*Operation, error)
	WriteOperations([]*Operation, string) error
}

type PackagesStorage interface {
	ReadPackages() ([]*Package, error)
	WritePackages([]*Package, string) error
}
