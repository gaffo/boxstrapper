package boxstrapper

import (
	"fmt"
	"sort"
	"strings"
)

func Ap(driver Driver, storage Storage, packages []string) error {
	// Load prevoius packages
	strPrexist, _ := storage.ReadPackages()
	mpPkgnamePkg := make(map[string]Package)
	aPackages := ParsePackages(strPrexist)
	for _, pkg := range aPackages {
		mpPkgnamePkg[pkg.Package] = pkg
	}

	// Install Packages
	for _, pkgStr := range packages {
		pkg := PackageFromApString(pkgStr)
		driver.AddPackage(pkg.Package)
		if oldPkg, ok := mpPkgnamePkg[pkg.Package]; ok {
			pkg.Groups = append(oldPkg.Groups, pkg.Groups...)
		}
		mpPkgnamePkg[pkg.Package] = pkg
	}

	// sort the packages
	aPackages = make([]Package, 0, len(mpPkgnamePkg))
	for _, pkg := range mpPkgnamePkg {
		aPackages = append(aPackages, pkg)
	}
	sort.Sort(ByPackageName(aPackages))

	// output the packages to file again
	sRet := ""
	for i, pkg := range aPackages {
		if i != 0 {
			sRet += "\n"
		}
		sRet += pkg.String()
	}
	storage.WritePackages(sRet, fmt.Sprintf("ap %s", strings.Join(packages, ", ")))

	return nil
}
