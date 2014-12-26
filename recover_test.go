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
	storage := new(mocks.OperationsStorage)
	storage.On("ReadOperations").Return([]*Operation{}, nil)

	err := Recover(driver, storage)
	assert.Nil(err)

	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func TestRecoverSinglePackage(t *testing.T) {
	assert := assert.New(t)
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package").Return(nil).Once()

	storage := new(mocks.OperationsStorage)
	storage.On("ReadOperations").Return(
		[]*Operation{
			OperationFromPackage(&Package{Name: "package", Groups: []string{"default"}}),
		}, nil)

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

	storage := new(mocks.OperationsStorage)
	storage.On("ReadOperations").Return(
		[]*Operation{
			OperationFromPackage(&Package{Name: "package", Groups: []string{"default"}}),
			OperationFromPackage(&Package{Name: "package2", Groups: []string{"dev"}}),
		}, nil)

	err := Recover(driver, storage)
	assert.Nil(err)

	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}
