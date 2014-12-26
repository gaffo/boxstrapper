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

type Operation struct {
	Name   string
	Params []string
	Groups []string
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

func (this Operation) String() string {
	this.Groups = removeDuplicates(this.Groups)
	sort.Strings(this.Groups)
	return fmt.Sprintf("%s(%s): %s", this.Name, strings.Join(this.Params, ", "), strings.Join(this.Groups, ", "))
}

type ByOperationName []*Operation

func (a ByOperationName) Len() int           { return len(a) }
func (a ByOperationName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOperationName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type Operations struct {
	Operations []Operation
}

func (this *Operations) Add(op Operation) {
	this.Operations = append(this.Operations, op)
}

func (this Operations) String() string {
	results := make([]string, len(this.Operations))
	for i, op := range this.Operations {
		results[i] = op.String()
	}
	return strings.Join(results, "\n")
}

func ParseOperations(contents string) []Operation {
	ops, err := ops_parser.ParseOps(contents)
	if err != nil {
		fmt.Println(err)
	}

	operations := make([]Operation, 0, len(ops))

	for _, op := range ops {
		operations = append(operations,
			Operation{
				Name:   op.Name,
				Params: op.Params,
				Groups: op.Groups})
	}
	return operations
}
