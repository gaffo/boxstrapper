package ops_parser_test

// tell go generate that we want to run  the command below.
//go:generate ragel -Z -G2 -o ops_parser.go ops_parser.rl
//go:generate ragel -Vp ops_parser.rl -o ops_parser.dot
//go:generate dot ops_parser.dot -Tpng -o ops_parser.png

import (
	"fmt"
	. "github.com/gaffo/boxstrapper/ops_parser"
	"testing"
)

func helper(input string, pkgs []Op, t *testing.T) {
	v, e := ParseOps(input)
	if e != nil {
		t.Fatalf("Errored on %s: %s", input, e)
	}
	if len(v) != len(pkgs) {
		t.Fatalf("Expected %d packages but was %d", len(pkgs), len(v))
	}
	for i, expected := range pkgs {
		actual := v[i]
		if expected.Name != actual.Name {
			t.Fatalf("%d: Expected name == '%s' but was '%s'", i, expected.Name, actual.Name)
		}

		if len(expected.Params) != len(actual.Params) {
			t.Fatalf("%d: Expected %d params but was %d", i, len(expected.Params), len(actual.Params))
		}
		for j, expectedParam := range expected.Params {
			actualParam := actual.Params[j]
			if expectedParam != actualParam {
				t.Fatalf("%d:%d Expected group %s but was %s", i, j, expectedParam, actualParam)
			}
		}

		if len(expected.Groups) != len(actual.Groups) {
			t.Fatalf("%d: Expected %d groups but was %d", i, len(expected.Groups), len(actual.Groups))
		}
		for j, expectedGroup := range expected.Groups {
			actualGroup := actual.Groups[j]
			if expectedGroup != actualGroup {
				t.Fatalf("%d:%d Expected group %s but was %s", i, j, expectedGroup, actualGroup)
			}
		}

		if len(expected.Depends) != len(actual.Depends) {
			t.Fatalf("%d: Expected %d Depends but was %d", i, len(expected.Depends), len(actual.Depends))
		}
		for j, expected := range expected.Depends {
			actual := actual.Depends[j]
			if expected.Operation != actual.Operation {
				t.Fatalf("%d:%d Expected operation %s but was %s", i, j, expected.Operation, actual.Operation)
			}

			if len(expected.Params) != len(actual.Params) {
				t.Fatalf("%d: Expected %d groups but was %d", i, len(expected.Params), len(actual.Params))
			}

			for k, expected := range expected.Params {
				actual := actual.Params[k]
				if expected != actual {
					t.Fatalf("%d:%d:%d Expected Param %s but was %s", i, j, k, expected, actual)
				}
			}
		}
	}
}

func errorTest(input, expected string, t *testing.T) {
	_, e := ParseOps(input)
	if e == nil {
		t.Fatalf("Expected parse error")
	}
	actual := fmt.Sprint(e)
	if expected != actual {
		t.Fatalf("Expected error [%s] but was [%s]", expected, actual)
	}
}

func Test_SinglePackage_SingleGroup(t *testing.T) {
	helper(
		"package(bar): baz",
		[]Op{Op{Name: "package", Params: []string{"bar"}, Groups: []string{"baz"}}},
		t)
}

func Test_SinglePackageNoSpace_SingleGroup(t *testing.T) {
	helper(
		"package(bar):baz",
		[]Op{Op{Name: "package", Params: []string{"bar"}, Groups: []string{"baz"}}},
		t)
}

func Test_SinglePackage_MultipleGroups(t *testing.T) {
	helper(
		"package(bar): baz, flork",
		[]Op{Op{Name: "package", Params: []string{"bar"}, Groups: []string{"baz", "flork"}}},
		t)
}

func Test_SinglePackage_MultipleGroups3(t *testing.T) {
	helper(
		"package(bar): baz, flork, flee",
		[]Op{Op{Name: "package", Params: []string{"bar"}, Groups: []string{"baz", "flork", "flee"}}},
		t)
}

func Test_SinglePackage_MultipleGroups3NoCommas(t *testing.T) {
	helper(
		"package(bar): baz, flork,flee",
		[]Op{Op{Name: "package", Params: []string{"bar"}, Groups: []string{"baz", "flork", "flee"}}},
		t)
}

func Test_MultiplePackages_OneGroup(t *testing.T) {
	helper(
		`package(bar): baz
package(flee): flicker`,
		[]Op{
			Op{Name: "package", Params: []string{"bar"}, Groups: []string{"baz"}},
			Op{Name: "package", Params: []string{"flee"}, Groups: []string{"flicker"}},
		},
		t)
}

func Test_MultiplePackages_MultipleGroups(t *testing.T) {
	helper(
		`package(bar): baz, snot
package(flee): flicker, statueoliberty, soccer`,
		[]Op{
			Op{Name: "package", Params: []string{"bar"}, Groups: []string{"baz", "snot"}},
			Op{Name: "package", Params: []string{"flee"}, Groups: []string{"flicker", "statueoliberty", "soccer"}},
		},
		t)
}

func Test_Error_OpNoColon(t *testing.T) {
	errorTest(
		"package(pkg)",
		`Parse error on line 1 char 13:
package(pkg)
------------^`, t)
}

func Test_Error_SecondLine(t *testing.T) {
	errorTest(
		`package(pkl): bob
package(bar)`,
		`Parse error on line 2 char 13:
package(bar)
------------^`, t)
}

func Test_Watch_SingleFile_SingleGroup(t *testing.T) {
	helper("watch(path/to/file): baz",
		[]Op{Op{Name: "watch", Params: []string{"path/to/file"}, Groups: []string{"baz"}}},
		t)
}

func Test_OperationMultipleParams(t *testing.T) {
	helper("watch(arg1, arg2, arg3): baz",
		[]Op{
			Op{
				Name:   "watch",
				Params: []string{"arg1", "arg2", "arg3"},
				Groups: []string{"baz"},
			},
		},
		t)
}

func Test_Dep_Single(t *testing.T) {
	helper("package(vim): development > package(gcc)",
		[]Op{
			Op{
				Name:   "package",
				Params: []string{"vim"},
				Groups: []string{"development"},
				Depends: []Dep{
					Dep{
						Operation: "package",
						Params:    []string{"gcc"},
					},
				},
			},
		},
		t)
}

func Test_Dep_MultipleDeps(t *testing.T) {
	helper("package(vim): development > package(gcc), call(arg1, arg2, arg3)",
		[]Op{
			Op{
				Name:   "package",
				Params: []string{"vim"},
				Groups: []string{"development"},
				Depends: []Dep{
					Dep{
						Operation: "package",
						Params:    []string{"gcc"},
					},
					Dep{
						Operation: "call",
						Params:    []string{"arg1", "arg2", "arg3"},
					},
				},
			},
		},
		t)
}

func Test_Dep_MultipleDeps_MultipleLines(t *testing.T) {
	helper(
		`package(vim): development > package(gcc), call(arg1, arg2, arg3)
package(emacs): default,development >package(gcc), watch(~/.zshrc)`,
		[]Op{
			Op{
				Name:   "package",
				Params: []string{"vim"},
				Groups: []string{"development"},
				Depends: []Dep{
					Dep{
						Operation: "package",
						Params:    []string{"gcc"},
					},
					Dep{
						Operation: "call",
						Params:    []string{"arg1", "arg2", "arg3"},
					},
				},
			},
			Op{
				Name:   "package",
				Params: []string{"emacs"},
				Groups: []string{"default", "development"},
				Depends: []Dep{
					Dep{
						Operation: "package",
						Params:    []string{"gcc"},
					},
					Dep{
						Operation: "watch",
						Params:    []string{"~/.zshrc"},
					},
				},
			},
		},
		t)
}

func strEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}
