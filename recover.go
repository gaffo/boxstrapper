package boxstrapper

import ()

func Recover(driver Driver, storage OperationsStorage) error {
	operations, _ := storage.ReadOperations()

	for _, op := range operations {
		pkg, _ := PackageFromOperation(op)
		driver.AddPackage(pkg.Name)
	}

	return nil
}
