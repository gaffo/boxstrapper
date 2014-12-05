package boxstrapper_test

import (
	"testing"
	. "github.com/gaffo/boxstrapper"
  	"github.com/stretchr/testify/assert"
  	"github.com/gaffo/boxstrapper/mocks"
)

func TestApCallsToDriver(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()

	storage := new(mocks.Storage)
	storage.On("WritePackages", "package1: default").Return(nil).Once()
	
	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func TestApCallsToDriver_MultiplePackages(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1", "package2"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()
	driver.On("AddPackage", "package2").Return(nil).Once()

	storage := new(mocks.Storage)
	storage.On("WritePackages", `package1: default
package2: default`).Return(nil).Once()
	
	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func TestApCallsToDriver_MultiplePackages_UnsortedPackages(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package2", "package1"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()
	driver.On("AddPackage", "package2").Return(nil).Once()

	storage := new(mocks.Storage)
	storage.On("WritePackages", `package1: default
package2: default`).Return(nil).Once()
	
	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func TestApCallsToDriver_WithSingleGroup(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1:system"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()

	storage := new(mocks.Storage)
	storage.On("WritePackages", "package1: system").Return(nil).Once()
	
	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func TestApCallsToDriver_MultiplePackages_WithSingleGroup(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1:george", "package2:system"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()
	driver.On("AddPackage", "package2").Return(nil).Once()

	storage := new(mocks.Storage)
	storage.On("WritePackages", `package1: george
package2: system`).Return(nil).Once()
	
	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func TestApCallsToDriver_MultiplePackages_UnsortedPackages_WithSingleGroup(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package2:system", "package1:george"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()
	driver.On("AddPackage", "package2").Return(nil).Once()

	storage := new(mocks.Storage)
	storage.On("WritePackages", `package1: george
package2: system`).Return(nil).Once()
	
	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}