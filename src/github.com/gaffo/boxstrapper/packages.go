package boxstrapper

import (
	"strings"
	// "fmt"
)

type Package struct {
	Package string
	Groups []string
}

func (this Package) String() string {
	return "hi: default"
}

type Packages struct {
}

func (this *Packages) Add(pkg Package) error {

	return nil
}

func ParsePackages(contents string) []Package {
	lines := strings.Split(contents, "\n")

	packages := make([]Package, 0, len(lines))

	for _, line := range(lines) {
		if line == "" {
			continue
		}

		if !strings.Contains(line, ": ") {
			line = strings.TrimSpace(line)
			packages = append(packages, Package{Package: line, Groups: []string{"default"}})
			continue
		}

		parts := strings.SplitN(line, ": ", 2)
		pkg := parts[0]
		pkg = strings.TrimSpace(pkg)
		groups := strings.Split(parts[1], ",")
		for i, group := range(groups) {
			groups[i] = strings.TrimSpace(group)
		}
		packages = append(packages, Package{Package: pkg, Groups: groups})
	}
	return packages
}
