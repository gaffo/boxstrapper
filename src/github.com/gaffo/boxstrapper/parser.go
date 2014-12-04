package boxstrapper

import (
	"strings"
	"fmt"
)

type Package struct {
	Package string
	Groups []string
}

func ParsePackages(contents string) []Package {
	lines := strings.Split(contents, "\n")
	fmt.Println(lines)

	packages := make([]Package, 0, len(lines))

	for _, line := range(lines) {
		fmt.Println(">", line)
		if line == "" {
			fmt.Println("Empty")
			continue
		}

		if !strings.Contains(line, ": ") {
			fmt.Println("Solo")
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
