package boxstrapper

import (
)

func Ap(driver Driver, storage Storage, packages []string) error {
	for _, pkg := range(packages) {
		driver.AddPackage(pkg)
	}
	storage.WritePackages("package1: default")
	return nil
}