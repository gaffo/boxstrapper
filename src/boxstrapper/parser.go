package boxstrapper

type Package struct {
	Package string
	Groups []string
}

func NewPackage(contents string) []Package {
	return []Package{Package{}}
}
