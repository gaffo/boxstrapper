package boxstrapper

import (
	"fmt"
	"sort"
	"strings"
)

func PackageFromApString(op string) *Package {
	parts := strings.Split(op, ":")
	if len(parts) == 1 {
		return &Package{Name: parts[0], Groups: []string{"default"}}
	}
	return &Package{Name: parts[0], Groups: strings.Split(parts[1], ",")}
}

func Ap(driver Driver, storage PackagesStorage, packages []string) error {
	// Load prevoius packages
	mpPkgnamePkg := make(map[string]*Package)
	aPackages, _ := storage.ReadPackages()
	for _, pkg := range aPackages {
		mpPkgnamePkg[pkg.Name] = pkg
	}

	// Install Operations
	aAddedPackages := make([]string, 0, len(packages))
	for _, pkgStr := range packages {
		pkg := PackageFromApString(pkgStr)
		driver.AddPackage(pkg.Name)
		if oldPkg, ok := mpPkgnamePkg[pkg.Name]; ok {
			oldPkg.Groups = append(oldPkg.Groups, pkg.Groups...)
		} else {
			aPackages = append(aPackages, pkg)
		}
		mpPkgnamePkg[pkg.Name] = pkg
		aAddedPackages = append(aAddedPackages, pkg.Name)
	}

	// get package names and sort them
	sort.Strings(aAddedPackages)

	// output the packages to file again
	storage.WritePackages(
		aPackages,
		fmt.Sprintf("added packages: %s", strings.Join(aAddedPackages, ", ")))

	return nil
}
