package boxstrapper

import ()

func Recover(driver Driver, storage Storage) error {
	strPackages, _ := storage.ReadPackages()

	for _, pkg := range ParsePackages(strPackages) {
		driver.AddPackage(pkg.Package)
	}

	return nil
}
