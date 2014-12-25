package boxstrapper_test

import (
	. "github.com/gaffo/boxstrapper"
	"github.com/gaffo/boxstrapper/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecoverNoPackages(t *testing.T) {
	assert := assert.New(t)
	driver := new(mocks.Driver)
	storage := new(mocks.Storage)
	storage.On("ReadPackages").Return("", nil)

	err := Recover(driver, storage)
	assert.Nil(err)

	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func TestRecoverSinglePackage(t *testing.T) {
	assert := assert.New(t)
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package").Return(nil).Once()

	storage := new(mocks.Storage)
	storage.On("ReadPackages").Return("package(package): default", nil)

	err := Recover(driver, storage)
	assert.Nil(err)

	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func TestRecoverMultiplePackages(t *testing.T) {
	assert := assert.New(t)
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package").Return(nil).Once()
	driver.On("AddPackage", "package2").Return(nil).Once()

	storage := new(mocks.Storage)
	storage.On("ReadPackages").Return(`package(package): default
package(package2): dev`, nil)

	err := Recover(driver, storage)
	assert.Nil(err)

	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}
