package boxstrapper

import (
	"strings"
)

type Package struct {
	Package string
	Groups []string
}

func NewPackage(contents string) []Package {
	if !strings.Contains(contents, ": ") {
		return []Package{Package{Package: contents, Groups: []string{"default"}}}
	}

	parts := strings.SplitN(contents, ": ", 2)
	pkg := parts[0]
	groups := strings.Split(parts[1], ",")
	for i, group := range(groups) {
		groups[i] = strings.TrimSpace(group)
	}

	return []Package{Package{Package: pkg, Groups: groups}}
}
