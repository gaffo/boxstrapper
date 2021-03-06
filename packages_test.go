package boxstrapper_test

import (
	. "github.com/gaffo/boxstrapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackageToString_SingleGroup(t *testing.T) {
	assert := assert.New(t)
	pkg := Package{Package: "hi", Groups: []string{"default"}}
	assert.Equal("package(hi): default", pkg.String())
}

func TestPackageToString_TwoGroups(t *testing.T) {
	assert := assert.New(t)
	pkg := Package{Package: "hi", Groups: []string{"default", "system"}}
	assert.Equal("package(hi): default, system", pkg.String())
}

func TestPackageToString_TwoGroupsAlphabetize(t *testing.T) {
	assert := assert.New(t)
	pkg := Package{Package: "hi", Groups: []string{"system", "default"}}
	assert.Equal("package(hi): default, system", pkg.String())
}

func TestPackageToString_ToGroupsDuplicate(t *testing.T) {
	assert := assert.New(t)
	pkg := Package{Package: "hi", Groups: []string{"system", "system"}}
	assert.Equal("package(hi): system", pkg.String())
}

func TestSinglePackage_NoGroups(t *testing.T) {
	assert := assert.New(t)
	s := "i3"

	packages := ParsePackages(s)

	assert.Equal(0, len(packages))
}

func TestSinglePackage_PrefixWhitespace_NoGroups(t *testing.T) {
	assert := assert.New(t)
	s := " i3"

	packages := ParsePackages(s)

	assert.Equal(0, len(packages))
}

func TestSinglePackage_PostfixWhitespace_NoGroups(t *testing.T) {
	assert := assert.New(t)
	s := "i3 "

	packages := ParsePackages(s)

	assert.Equal(0, len(packages))
}

func TestSinglePackageWithLeadingWhitespace(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3):  system"

	packages := ParsePackages(s)

	assert.Equal(1, len(packages))
	assert.Equal("i3", packages[0].Package)
	assert.Equal(1, len(packages[0].Groups))
	assert.Equal("system", packages[0].Groups[0])
}

func TestSinglePackage_PrefixWhitespace_WithGroup(t *testing.T) {
	assert := assert.New(t)
	s := " packages(i3): system"

	packages := ParsePackages(s)

	assert.Equal(1, len(packages))
	assert.Equal("i3", packages[0].Package)
	assert.Equal(1, len(packages[0].Groups))
	assert.Equal("system", packages[0].Groups[0])
}

func TestSinglePackage_PostfixWhitespace_WithGroup(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3) : system"

	packages := ParsePackages(s)

	assert.Equal(0, len(packages))
}

func TestSinglePackageWithTrailingWhitespace(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3): system "

	packages := ParsePackages(s)

	assert.Equal(0, len(packages))
}

func TestSinglePackage_MultipleGroups(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3): system, boxstrapper"

	packages := ParsePackages(s)

	assert.Equal(1, len(packages))
	assert.Equal("i3", packages[0].Package)
	assert.Equal(2, len(packages[0].Groups))
	assert.Equal("system", packages[0].Groups[0])
	assert.Equal("boxstrapper", packages[0].Groups[1])
}

func TestSinglePackage_MultipleGroups_PrefixWhitespace(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3): system,  boxstrapper"

	packages := ParsePackages(s)

	assert.Equal(0, len(packages))
}

func TestSinglePackage_MultipleGroups_PostfixWhitespace(t *testing.T) {
	assert := assert.New(t)
	s := "package(i3): system, boxstrapper "

	packages := ParsePackages(s)

	assert.Equal(0, len(packages))
}

func TestMultiplePackages_SingleGroups(t *testing.T) {
	assert := assert.New(t)
	s := `package(i3): system
package(boxstrapper): dev`

	packages := ParsePackages(s)

	assert.Equal(2, len(packages))

	assert.Equal("i3", packages[0].Package)
	assert.Equal(1, len(packages[0].Groups))
	assert.Equal("system", packages[0].Groups[0])

	assert.Equal("boxstrapper", packages[1].Package)
	assert.Equal(1, len(packages[1].Groups))
	assert.Equal("dev", packages[1].Groups[0])
}

func TestMultiplePackages_MultipleGroups(t *testing.T) {
	assert := assert.New(t)
	s := `package(i3): system, long
package(boxstrapper): dev, pork`

	packages := ParsePackages(s)

	assert.Equal(2, len(packages))

	assert.Equal("i3", packages[0].Package)
	assert.Equal(2, len(packages[0].Groups))
	assert.Equal("system", packages[0].Groups[0])
	assert.Equal("long", packages[0].Groups[1])

	assert.Equal("boxstrapper", packages[1].Package)
	assert.Equal(2, len(packages[1].Groups))
	assert.Equal("dev", packages[1].Groups[0])
	assert.Equal("pork", packages[1].Groups[1])
}

func TestMultiplePackages_EmptyThirdPackage(t *testing.T) {
	assert := assert.New(t)
	s := `package(i3): system, long
package(boxstrapper): dev, pork
`

	packages := ParsePackages(s)

	assert.Equal(0, len(packages))
}

func TestPackages_Print(t *testing.T) {
	assert := assert.New(t)
	packages := Packages{}
	packages.Add(Package{Package: "pkg", Groups: []string{"g1", "g2"}})

	result := packages.String()

	assert.Equal("package(pkg): g1, g2", result)
}

func TestPackages_MultiplePackages(t *testing.T) {
	assert := assert.New(t)
	packages := Packages{}
	packages.Add(Package{Package: "pkg", Groups: []string{"g1", "g2"}})
	packages.Add(Package{Package: "pkg2", Groups: []string{"g4", "g3"}})

	result := packages.String()

	assert.Equal(`package(pkg): g1, g2
package(pkg2): g3, g4`, result)
}

func TestPackageFromApString_NoGroup(t *testing.T) {
	assert := assert.New(t)

	pkg := PackageFromApString("package")
	assert.Equal("package", pkg.Package)
	assert.Equal(1, len(pkg.Groups))
	assert.Equal("default", pkg.Groups[0])
}

func TestPackageFromApString_SingleGroup(t *testing.T) {
	assert := assert.New(t)

	pkg := PackageFromApString("package:system")
	assert.Equal("package", pkg.Package)
	assert.Equal(1, len(pkg.Groups))
	assert.Equal("system", pkg.Groups[0])
}

func TestPackageFromApString_MultipleGroup(t *testing.T) {
	assert := assert.New(t)

	pkg := PackageFromApString("package:system,george")
	assert.Equal("package", pkg.Package)
	assert.Equal(2, len(pkg.Groups))
	assert.Equal("system", pkg.Groups[0])
	assert.Equal("george", pkg.Groups[1])
}
