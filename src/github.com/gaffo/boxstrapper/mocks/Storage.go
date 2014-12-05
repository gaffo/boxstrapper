package mocks

import "github.com/stretchr/testify/mock"

type Storage struct {
	mock.Mock
}

func (m *Storage) ReadPackages() (string, error) {
	ret := m.Called()

	r0 := ret.Get(0).(string)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *Storage) WritePackages(contents string) error {
	ret := m.Called(contents)

	r0 := ret.Error(0)

	return r0
}
