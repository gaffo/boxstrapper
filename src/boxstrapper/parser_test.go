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
}