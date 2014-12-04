package boxstraper

import (
	"fmt"
)

type UbuntuDriver struct {
}

func (*UbuntuDriver) AddPackage(packageName string) err {
	fmt.Println("Installing Package:", packageName)
	return nil
}