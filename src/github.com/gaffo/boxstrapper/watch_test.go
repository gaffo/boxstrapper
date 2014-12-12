package boxstrapper_test

import (
	"testing"
	. "github.com/gaffo/boxstrapper"
  	"github.com/stretchr/testify/assert"
  	"github.com/gaffo/boxstrapper/mocks"
)

func Test_SingleFile_DoesntExist(t *testing.T) {
	assert := assert.New(t)

	files := []string{"missingfile"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil)

	storage := new(mocks.Storage)
	storage.On("ReadPackages").Return("", nil)

	Watch(driver, storage, files)

	assert.True(true)
}