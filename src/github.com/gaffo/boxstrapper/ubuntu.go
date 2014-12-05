package boxstrapper

import (
	"fmt"
)

type UbuntuDriver struct {
}

func (UbuntuDriver) AddPackage(packageName string) error {
	fmt.Println("Installing Package:", packageName)
	return nil
}