package boxstrapper

import (
	// "os"
)

type Driver interface {
	AddPackage(packageName string) error
}