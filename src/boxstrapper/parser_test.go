package boxstrapper_test

import (
	"testing"
	. "boxstrapper"
  	"github.com/stretchr/testify/assert"
)

func TestSinglePackage(t *testing.T) {
	s := "i3: system"

	packages := NewPackage(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
}

func TestSinglePackageWithLeadingWhitespace(t *testing.T) {
	s := "i3:  system"

	packages := NewPackage(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
}

func TestSinglePackageWithTrailingWhitespace(t *testing.T) {
	s := "i3: system "

	packages := NewPackage(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 1, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
}

func TestSinglePackage_MultipleGroups(t *testing.T) {
	s := "i3: system, boxstrapper"

	packages := NewPackage(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 2, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
	assert.Equal(t, "boxstrapper", packages[0].Groups[1])
}
	
func TestSinglePackage_MultipleGroups_PrefixWhitespace(t *testing.T) {
	s := "i3: system,  boxstrapper"

	packages := NewPackage(s)

	assert.Equal(t, 1, len(packages))
	assert.Equal(t, "i3", packages[0].Package)
	assert.Equal(t, 2, len(packages[0].Groups))
	assert.Equal(t, "system", packages[0].Groups[0])
	assert.Equal(t, "boxstrapper", packages[0].Groups[1])
}