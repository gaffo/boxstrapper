package boxstrapper

type FilesystemStorage struct {
}

func (FilesystemStorage) ReadPackages() (string, error) {
	return "", nil
}

func (FilesystemStorage) WritePackages(contents string) error {
	return nil
}