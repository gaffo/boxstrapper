package boxstrapper_test

import (
	. "github.com/gaffo/boxstrapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_OperationToString_SingleGroup(t *testing.T) {
	assert := assert.New(t)
	pkg := Operation{Name: "hi", Groups: []string{"default"}}
	assert.Equal("hi(): default", pkg.String())
}

func Test_OperationToString_TwoGroups(t *testing.T) {
	assert := assert.New(t)
	pkg := Operation{Name: "hi", Params: []string{"p1", "p2"}, Groups: []string{"default", "system"}}
	assert.Equal("hi(p1, p2): default, system", pkg.String())
}

func Test_OperationToString_TwoGroupsAlphabetize(t *testing.T) {
	assert := assert.New(t)
	pkg := Operation{Name: "hi", Groups: []string{"system", "default"}}
	assert.Equal("hi(): default, system", pkg.String())
}

func Test_OperationToString_ToGroupsDuplicate(t *testing.T) {
	assert := assert.New(t)
	pkg := Operation{Name: "hi", Groups: []string{"system", "system"}}
	assert.Equal("hi(): system", pkg.String())
}

func TestSinglePackage_NoGroups(t *testing.T) {
	assert := assert.New(t)
	s := "i3"

	operations := ParseOperations(s)

	assert.Equal(0, len(operations))
}

func TestSinglePackage_PrefixWhitespace_NoGroups(t *testing.T) {
	assert := assert.New(t)
	s := " i3"

	operations := ParseOperations(s)

	assert.Equal(0, len(operations))
}

func TestSinglePackage_PostfixWhitespace_NoGroups(t *testing.T) {
	assert := assert.New(t)
	s := "i3 "

	operations := ParseOperations(s)

	assert.Equal(0, len(operations))
}

func TestSinglePackageWithLeadingWhitespace(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3):  system"

	operations := ParseOperations(s)

	assert.Equal(1, len(operations))
	assert.Equal("i3", operations[0].Params[0])
	assert.Equal(1, len(operations[0].Groups))
	assert.Equal("system", operations[0].Groups[0])
}

func TestSinglePackage_PrefixWhitespace_WithGroup(t *testing.T) {
	assert := assert.New(t)
	s := " operations(i3): system"

	operations := ParseOperations(s)

	assert.Equal(1, len(operations))
	assert.Equal("i3", operations[0].Params[0])
	assert.Equal(1, len(operations[0].Groups))
	assert.Equal("system", operations[0].Groups[0])
}

func TestSinglePackage_PostfixWhitespace_WithGroup(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3) : system"

	operations := ParseOperations(s)

	assert.Equal(0, len(operations))
}

func TestSinglePackageWithTrailingWhitespace(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3): system "

	operations := ParseOperations(s)

	assert.Equal(0, len(operations))
}

func TestSinglePackage_MultipleGroups(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3): system, boxstrapper"

	operations := ParseOperations(s)

	assert.Equal(1, len(operations))
	assert.Equal("i3", operations[0].Params[0])
	assert.Equal(2, len(operations[0].Groups))
	assert.Equal("system", operations[0].Groups[0])
	assert.Equal("boxstrapper", operations[0].Groups[1])
}

func TestSinglePackage_MultipleGroups_PrefixWhitespace(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3): system,  boxstrapper"

	operations := ParseOperations(s)

	assert.Equal(0, len(operations))
}

func TestSinglePackage_MultipleGroups_PostfixWhitespace(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3): system, boxstrapper "

	operations := ParseOperations(s)

	assert.Equal(0, len(operations))
}

func TestMultiplePackages_SingleGroups(t *testing.T) {
	assert := assert.New(t)
	s := `package(i3): system
package(boxstrapper): dev`

	operations := ParseOperations(s)

	assert.Equal(2, len(operations))

	assert.Equal("i3", operations[0].Params[0])
	assert.Equal(1, len(operations[0].Groups))
	assert.Equal("system", operations[0].Groups[0])

	assert.Equal("boxstrapper", operations[1].Params[0])
	assert.Equal(1, len(operations[1].Groups))
	assert.Equal("dev", operations[1].Groups[0])
}

func TestMultiplePackages_MultipleGroups(t *testing.T) {
	assert := assert.New(t)
	s := `package(i3): system, long
package(boxstrapper): dev, pork`

	operations := ParseOperations(s)

	assert.Equal(2, len(operations))

	assert.Equal("i3", operations[0].Params[0])
	assert.Equal(2, len(operations[0].Groups))
	assert.Equal("system", operations[0].Groups[0])
	assert.Equal("long", operations[0].Groups[1])

	assert.Equal("boxstrapper", operations[1].Params[0])
	assert.Equal(2, len(operations[1].Groups))
	assert.Equal("dev", operations[1].Groups[0])
	assert.Equal("pork", operations[1].Groups[1])
}

func TestMultiplePackages_EmptyThirdPackage(t *testing.T) {
	assert := assert.New(t)
	s := `package(i3): system, long
package(boxstrapper): dev, pork
`

	operations := ParseOperations(s)

	assert.Equal(0, len(operations))
}

func Test_Operations_Print(t *testing.T) {
	assert := assert.New(t)
	operations := Operations{}
	operations.Add(Operation{Name: "package", Params: []string{"pkg"}, Groups: []string{"g1", "g2"}})

	result := operations.String()

	assert.Equal("package(pkg): g1, g2", result)
}

func Test_Operations_MultiplePackages(t *testing.T) {
	assert := assert.New(t)
	operations := Operations{}
	operations.Add(Operation{Name: "package", Params: []string{"pkg"}, Groups: []string{"g1", "g2"}})
	operations.Add(Operation{Name: "package", Params: []string{"pkg2"}, Groups: []string{"g4", "g3"}})

	result := operations.String()

	assert.Equal(`package(pkg): g1, g2
package(pkg2): g3, g4`, result)
}
