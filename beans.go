package boxstrapper

import (
	"errors"
	"fmt"
)

type Package struct {
	Name   string
	Groups []string
}

func PackageFromOperation(operation *Operation) (*Package, error) {
	if operation == nil {
		return nil, errors.New("Can't convert nil operation to package")
	}

	if operation.Name != "package" {
		return nil, errors.New(fmt.Sprintf("Trying to create a package off of %s", operation))
	}

	if len(operation.Params) != 1 {
		return nil, errors.New(fmt.Sprintf("Creating packge from operation with %d args (!= 1)", len(operation.Params)))
	}

	return &Package{
		Name:   operation.Params[0],
		Groups: operation.Groups,
	}, nil
}

func OperationFromPackage(pkg *Package) *Operation {
	return &Operation{
		Name:   "package",
		Params: []string{pkg.Name},
		Groups: pkg.Groups,
	}
}
