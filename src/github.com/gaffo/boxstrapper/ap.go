package boxstrapper

import (
)

func Ap(driver Driver, packages []string) error {
	for _, pkg := range(packages) {
		driver.AddPackage(pkg)
	}
	return nil
}