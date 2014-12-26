package mocks

import "github.com/gaffo/boxstrapper"
import "github.com/stretchr/testify/mock"

type OperationsStorage struct {
	mock.Mock
}

func (m *OperationsStorage) ReadOperations() ([]*boxstrapper.Operation, error) {
	ret := m.Called()

	r0 := ret.Get(0).([]*boxstrapper.Operation)
	r1 := ret.Error(1)

	return r0, r1
}
func (m *OperationsStorage) WriteOperations(_a0 []*boxstrapper.Operation, _a1 string) error {
	ret := m.Called(_a0, _a1)

	r0 := ret.Error(0)

	return r0
}
