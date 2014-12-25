package boxstrapper_test

import (
	. "github.com/gaffo/boxstrapper"
	"github.com/gaffo/boxstrapper/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApCallsToDriver(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil)

	storage := new(mocks.Storage)
	storage.On("ReadPackages").Return("", nil)
	storage.On("WritePackages", "package(package1): default").Return(nil).Once()

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
	storage.On("ReadPackages").Return("", nil)
	storage.On("WritePackages", `package(package1): default
package(package2): default`).Return(nil).Once()

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
	storage.On("ReadPackages").Return("", nil)
	storage.On("WritePackages", `package(package1): default
package(package2): default`).Return(nil).Once()

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
	storage.On("ReadPackages").Return("", nil)
	storage.On("WritePackages", "package(package1): system").Return(nil).Once()

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
	storage.On("ReadPackages").Return("", nil)
	storage.On("WritePackages", `package(package1): george
package(package2): system`).Return(nil).Once()

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
	storage.On("ReadPackages").Return("", nil)
	storage.On("WritePackages", `package(package1): george
package(package2): system`).Return(nil).Once()

	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func TestAp_WithDifferentPreexisting_DoesntLosePrexisting(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()

	storage := new(mocks.Storage)
	storage.On("ReadPackages").Return("package(package2): default", nil)
	storage.On("WritePackages", `package(package1): default
package(package2): default`).Return(nil).Once()

	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func TestAp_PreExistingGroup_IsntAddedToFile(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()

	storage := new(mocks.Storage)
	storage.On("ReadPackages").Return("package1", nil)
	storage.On("WritePackages", `package(package1): default`).Return(nil).Once()

	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func TestAp_PreExistingPackage_NewGroup_MergesGroups(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1:system"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()

	storage := new(mocks.Storage)
	storage.On("ReadPackages").Return("package(package1): default", nil)
	storage.On("WritePackages", `package(package1): default, system`).Return(nil).Once()

	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}
