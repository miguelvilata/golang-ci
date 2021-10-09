package library

import (
	"github.com/stretchr/testify/mock"
)

// RepositoryMock Mysql Repostiroy
type RepositoryMock struct {
	mock.Mock
}

// Save mock
func (l *RepositoryMock) Save(book Book) (int64, error) {
	return 0, nil
}

// Delete mock
func (l *RepositoryMock) Delete(id int64) error {
	args := l.Called(id)
	return args.Error(0)
}

// Find mock
func (l *RepositoryMock) Find(id int64) (*Book, error) {
	args := l.Called(id)
	return args.Get(0).(*Book), nil
}

// Update mock
func (l *RepositoryMock) Update(id int64, book Book) (int64, error) {
	args := l.Called(id, book)
	return args.Get(0).(int64), args.Error(1)
}

// All mock
func (l *RepositoryMock) All() ([]Book, error) {
	args := l.Called()
	return args.Get(0).([]Book), nil
}
