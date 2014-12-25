package boxstrapper

import (
	"fmt"
	"github.com/gaffo/boxstrapper/ops_parser"
	"sort"
	"strings"
)

// tell go generate that we want to run  the command below.
//go:generate ragel -Z -G2 -o ops_parser/ops_parser.go ops_parser/ops_parser.rl
//go:generate ragel -Vp ops_parser/ops_parser.rl -o ops_parser/ops_parser.dot
//go:generate dot ops_parser/ops_parser.dot -Tpng -o ops_parser/ops_parser.png

type Package struct {
	Package string
	Groups  []string
}

func removeDuplicates(a []string) []string {
	result := []string{}
	seen := map[string]string{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

func (this Package) String() string {
	this.Groups = removeDuplicates(this.Groups)
	sort.Strings(this.Groups)
	return fmt.Sprintf("package(%s): %s", this.Package, strings.Join(this.Groups, ", "))
}

func PackageFromApString(pkg string) Package {
	parts := strings.Split(pkg, ":")
	if len(parts) == 1 {
		return Package{Package: parts[0], Groups: []string{"default"}}
	}
	return Package{Package: parts[0], Groups: strings.Split(parts[1], ",")}
}

type ByPackageName []Package

func (a ByPackageName) Len() int           { return len(a) }
func (a ByPackageName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPackageName) Less(i, j int) bool { return a[i].Package < a[j].Package }

type Packages struct {
	Packages []Package
}

func (this *Packages) Add(pkg Package) {
	this.Packages = append(this.Packages, pkg)
}

func (this Packages) String() string {
	results := make([]string, len(this.Packages))
	for i, pkg := range this.Packages {
		results[i] = pkg.String()
	}
	return strings.Join(results, "\n")
}

func ParsePackages(contents string) []Package {
	ops, _ := ops_parser.ParseOps(contents)

	packages := make([]Package, 0, len(ops))

	for _, op := range ops {
		packages = append(packages,
			Package{
				Package: op.Params[0],
				Groups:  op.Groups})
	}
	return packages
}
