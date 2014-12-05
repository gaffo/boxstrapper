package boxstrapper

import (
	"strings"
	"sort"
)

func Ap(driver Driver, storage Storage, packages []string) error {
	pkgContents := make([]string, 0, len(packages))
	for _, pkg := range(packages) {
		driver.AddPackage(pkg)
		sPkg := Package{Package: pkg, Groups: []string{"default"}}.String()
		pkgContents = append(pkgContents, sPkg)
	}

	sort.Strings(pkgContents)
	sPkgContents := strings.Join(pkgContents, "\n")
	storage.WritePackages(sPkgContents)
	return nil
}