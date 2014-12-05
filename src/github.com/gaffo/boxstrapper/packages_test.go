package boxstrapper_test

import (
	"testing"
	. "github.com/gaffo/boxstrapper"
  	"github.com/stretchr/testify/assert"
)

func TestPackageToString_SingleGroup(t *testing.T) {
	assert := assert.New(t)
	pkg := Package{Package: "hi", Groups: []string{"default"}}
	assert.Equal("hi: default", pkg.String())
}

func TestSinglePackage_NoGroups(t *testing.T) {
	s := "i3"

	packages := ParsePackages(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "default", packages[0].Groups[0])
}

func TestSinglePackage_PrefixWhitespace_NoGroups(t *testing.T) {
	s := " i3"

	packages := ParsePackages(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "default", packages[0].Groups[0])
}

func TestSinglePackage_PostfixWhitespace_NoGroups(t *testing.T) {
	s := "i3 "

	packages := ParsePackages(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "default", packages[0].Groups[0])
}

func TestSinglePackageWithLeadingWhitespace(t *testing.T) {
	s := "i3:  system"

	packages := ParsePackages(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
}

func TestSinglePackage_PrefixWhitespace_WithGroup(t *testing.T) {
	s := " i3: system"

	packages := ParsePackages(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
}

func TestSinglePackage_PostfixWhitespace_WithGroup(t *testing.T) {
	s := "i3 : system"

	packages := ParsePackages(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
}

func TestSinglePackageWithTrailingWhitespace(t *testing.T) {
	s := "i3: system "

	packages := ParsePackages(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
}

func TestSinglePackage_MultipleGroups(t *testing.T) {
	s := "i3: system, boxstrapper"

	packages := ParsePackages(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 2, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
	assert.Equal(t, "boxstrapper", packages[0].Groups[1])
}
	
func TestSinglePackage_MultipleGroups_PrefixWhitespace(t *testing.T) {
	s := "i3: system,  boxstrapper"

	packages := ParsePackages(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 2, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
	assert.Equal(t, "boxstrapper", packages[0].Groups[1])
}
	
func TestSinglePackage_MultipleGroups_PostfixWhitespace(t *testing.T) {
	s := "i3: system, boxstrapper "

	packages := ParsePackages(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 2, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
	assert.Equal(t, "boxstrapper", packages[0].Groups[1])
}

func TestMultiplePackages_DefaultGroups(t *testing.T) {
	s := `i3
boxstrapper`

	packages := ParsePackages(s)

	assert.Equal(t, 2, len(packages))

	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "default", packages[0].Groups[0])

	assert.Equal(t, "boxstrapper", packages[1].Package)
	assert.Equal(t, 1, len(packages[1].Groups))
	assert.Equal(t, "default", packages[1].Groups[0])
}

func TestMultiplePackages_SingleGroups(t *testing.T) {
	s := `i3: system
boxstrapper: dev`

	packages := ParsePackages(s)

	assert.Equal(t, 2, len(packages))

	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])

	assert.Equal(t, "boxstrapper", packages[1].Package)
	assert.Equal(t, 1, len(packages[1].Groups))
	assert.Equal(t, "dev", packages[1].Groups[0])
}

func TestMultiplePackages_MultipleGroups(t *testing.T) {
	s := `i3: system, long
boxstrapper: dev, pork`

	packages := ParsePackages(s)

	assert.Equal(t, 2, len(packages))

	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 2, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
	assert.Equal(t, "long", packages[0].Groups[1])

	assert.Equal(t, "boxstrapper", packages[1].Package)
	assert.Equal(t, 2, len(packages[1].Groups))
	assert.Equal(t, "dev", packages[1].Groups[0])
	assert.Equal(t, "pork", packages[1].Groups[1])
}

func TestMultiplePackages_EmptyThirdPackage(t *testing.T) {
	s := `i3: system, long
boxstrapper: dev, pork
`

	packages := ParsePackages(s)

	assert.Equal(t, 2, len(packages))

	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 2, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
	assert.Equal(t, "long", packages[0].Groups[1])

	assert.Equal(t, "boxstrapper", packages[1].Package)
	assert.Equal(t, 2, len(packages[1].Groups))
	assert.Equal(t, "dev", packages[1].Groups[0])
	assert.Equal(t, "pork", packages[1].Groups[1])
}