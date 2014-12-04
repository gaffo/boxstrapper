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