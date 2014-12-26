package boxstrapper_test

//go:generate mockery --all github.com/gaffo/boxstrapper

import (
	. "github.com/gaffo/boxstrapper"
	"github.com/gaffo/boxstrapper/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockPackageStorage struct {
	readPackages    []*Package
	writtenPackages []*Package
	reason          string
	writeCount      int
}

func (this *MockPackageStorage) ReadPackages() ([]*Package, error) {
	return this.readPackages, nil
}

func (this *MockPackageStorage) WritePackages(packages []*Package, reason string) error {
	this.writeCount += 1
	this.reason = reason
	this.writtenPackages = packages
	return nil
}

func Test_Ap_NoPrevious_SinglePackage(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil)

	storage := &MockPackageStorage{writeCount: 0, readPackages: []*Package{}}

	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)

	assert.Equal("added packages: package1", storage.reason)
	assert.Equal(1, storage.writeCount)
	assert.Equal(1, len(storage.writtenPackages))
	assert.Equal("package1", storage.writtenPackages[0].Name)
	assert.Equal(1, len(storage.writtenPackages[0].Groups))
	assert.Equal("default", storage.writtenPackages[0].Groups[0])
}

func Test_Ap_NoPrevious_MultiplePackages(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1", "package2"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()
	driver.On("AddPackage", "package2").Return(nil).Once()

	storage := new(mocks.PackagesStorage)
	storage.On("ReadPackages").Return([]*Package{}, nil)
	storage.On("WritePackages",
		[]*Package{
			&Package{Name: "package1", Groups: []string{"default"}},
			&Package{Name: "package2", Groups: []string{"default"}},
		},
		"added packages: package1, package2").Return(nil).Once()

	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	storage.Mock.AssertExpectations(t)
}

func Test_Ap_NoPrevious_WithSingleGroup(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1:system"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()

	storage := &MockPackageStorage{writeCount: 0, readPackages: []*Package{}}

	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)
	assert.Equal("added packages: package1", storage.reason)
	assert.Equal(1, storage.writeCount)
	assert.Equal(1, len(storage.writtenPackages))
	assert.Equal("package1", storage.writtenPackages[0].Name)
	assert.Equal(1, len(storage.writtenPackages[0].Groups))
	assert.Equal("system", storage.writtenPackages[0].Groups[0])
}

func Test_Ap_NoPrevious_MultiplePackages_WithSingleGroup(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1:george", "package2:system"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()
	driver.On("AddPackage", "package2").Return(nil).Once()

	storage := &MockPackageStorage{writeCount: 0, readPackages: []*Package{}}

	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)

	assert.Equal("added packages: package1, package2", storage.reason)
	assert.Equal(1, storage.writeCount)
	assert.Equal(2, len(storage.writtenPackages))
	assert.Equal("package1", storage.writtenPackages[0].Name)
	assert.Equal(1, len(storage.writtenPackages[0].Groups))
	assert.Equal("george", storage.writtenPackages[0].Groups[0])

	assert.Equal("package2", storage.writtenPackages[1].Name)
	assert.Equal(1, len(storage.writtenPackages[1].Groups))
	assert.Equal("system", storage.writtenPackages[1].Groups[0])
}

func Test_Ap_WithPreviousNotSame_DoesntLosePrexisting(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()

	storage := &MockPackageStorage{
		writeCount: 0,
		readPackages: []*Package{
			&Package{Name: "package2", Groups: []string{"default"}},
		}}

	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)

	assert.Equal("added packages: package1", storage.reason)
	assert.Equal(1, storage.writeCount)
	assert.Equal(2, len(storage.writtenPackages))
	assert.Equal("package2", storage.writtenPackages[0].Name)
	assert.Equal(1, len(storage.writtenPackages[0].Groups))
	assert.Equal("default", storage.writtenPackages[0].Groups[0])

	assert.Equal("package1", storage.writtenPackages[1].Name)
	assert.Equal(1, len(storage.writtenPackages[1].Groups))
	assert.Equal("default", storage.writtenPackages[1].Groups[0])
}

func TestAp_PreExistingGroup_IsntAddedToFile(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()

	storage := &MockPackageStorage{
		writeCount: 0,
		readPackages: []*Package{
			&Package{Name: "package2", Groups: []string{"default"}},
		}}

	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)

	assert.Equal("added packages: package1", storage.reason)
	assert.Equal(1, storage.writeCount)
	assert.Equal(2, len(storage.writtenPackages))
	assert.Equal("package2", storage.writtenPackages[0].Name)
	assert.Equal(1, len(storage.writtenPackages[0].Groups))
	assert.Equal("default", storage.writtenPackages[0].Groups[0])

	assert.Equal("package1", storage.writtenPackages[1].Name)
	assert.Equal(1, len(storage.writtenPackages[1].Groups))
	assert.Equal("default", storage.writtenPackages[1].Groups[0])
}

func TestAp_PreExistingPackage_NewGroup_MergesGroups(t *testing.T) {
	assert := assert.New(t)

	packages := []string{"package1:system"}
	driver := new(mocks.Driver)
	driver.On("AddPackage", "package1").Return(nil).Once()

	storage := &MockPackageStorage{
		writeCount: 0,
		readPackages: []*Package{
			&Package{Name: "package1", Groups: []string{"default"}},
		}}

	err := Ap(driver, storage, packages)

	assert.Nil(err)
	driver.Mock.AssertExpectations(t)

	assert.Equal("added packages: package1", storage.reason)
	assert.Equal(1, storage.writeCount)
	assert.Equal(1, len(storage.writtenPackages))
	assert.Equal("package1", storage.writtenPackages[0].Name)
	assert.Equal(2, len(storage.writtenPackages[0].Groups))
	assert.Equal("default", storage.writtenPackages[0].Groups[0])
	assert.Equal("system", storage.writtenPackages[0].Groups[1])
}

func Test_OperationFromApString_NoGroup(t *testing.T) {
	assert := assert.New(t)

	pkg := PackageFromApString("package")
	assert.Equal("package", pkg.Name)
	assert.Equal(1, len(pkg.Groups))
	assert.Equal("default", pkg.Groups[0])
}

func Test_OperationFromApString_SingleGroup(t *testing.T) {
	assert := assert.New(t)

	pkg := PackageFromApString("package:system")
	assert.Equal("package", pkg.Name)
	assert.Equal(1, len(pkg.Groups))
	assert.Equal("system", pkg.Groups[0])
}

func Test_OperationFromApString_MultipleGroup(t *testing.T) {
	assert := assert.New(t)

	pkg := PackageFromApString("package:system,george")
	assert.Equal("package", pkg.Name)
	assert.Equal(2, len(pkg.Groups))
	assert.Equal("system", pkg.Groups[0])
	assert.Equal("george", pkg.Groups[1])
}
