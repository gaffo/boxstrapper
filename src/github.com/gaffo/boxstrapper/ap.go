package boxstrapper

import (
	"sort"
	"strings"
	"fmt"
)

func Ap(driver Driver, storage Storage, packages []string) error {
	pkgContents := make([]string, 0, len(packages))
	for _, pkgStr := range(packages) {
		pkg := PackageFromApString(pkgStr)
		driver.AddPackage(pkg.Package)
		sPkg := pkg.String()
		pkgContents = append(pkgContents, sPkg)
	}

	sort.Strings(pkgContents)
	fmt.Println(">", pkgContents)
	sPkgContents := strings.Join(pkgContents, "\n")
	storage.WritePackages(sPkgContents)
	return nil
}