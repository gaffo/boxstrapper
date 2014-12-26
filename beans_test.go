package boxstrapper_test

import (
	. "github.com/gaffo/boxstrapper"
	// "github.com/gaffo/boxstrapper/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PackageFromOperation_Nil(t *testing.T) {
	assert := assert.New(t)

	p, e := PackageFromOperation(nil)

	assert.Nil(p)
	assert.NotNil(e)
}

func Test_PackageFromOperation_Correct(t *testing.T) {
	assert := assert.New(t)

	p, e := PackageFromOperation(&Operation{Name: "package", Params: []string{"pkg"}, Groups: []string{"g1", "g2"}})

	assert.Nil(e)
	assert.Equal("pkg", p.Name)
	assert.Equal(2, len(p.Groups))
	assert.Equal("g1", p.Groups[0])
	assert.Equal("g2", p.Groups[1])
}
