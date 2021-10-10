package library

import (
	"testing"
)

func TestFind(t *testing.T) {
	repoMock := &RepositoryMock{}
	library := NewLibrary(repoMock)

	book := &Book{}
	var id int64 = 1
	repoMock.On("Find", id).Return(book)

	library.Find(1)

	repoMock.AssertNumberOfCalls(t, "Find", 1)
	repoMock.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	repoMock := &RepositoryMock{}
	library := NewLibrary(repoMock)

	var id int64 = 1
	repoMock.On("Delete", id).Return(nil)

	library.Delete(1)

	repoMock.AssertNumberOfCalls(t, "Delete", 1)
	repoMock.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	repoMock := &RepositoryMock{}
	library := NewLibrary(repoMock)

	var id int64 = 1
	book := &Book{}
	repoMock.On("Update", id, *book).Return(id, nil)

	library.Modify(id, *book)

	repoMock.AssertNumberOfCalls(t, "Update", 1)
	repoMock.AssertExpectations(t)
}

func TestAll(t *testing.T) {
	repoMock := &RepositoryMock{}
	library := NewLibrary(repoMock)

	books := []Book{}
	repoMock.On("All").Return(books, nil)

	library.All()

	repoMock.AssertNumberOfCalls(t, "All", 1)
	repoMock.AssertExpectations(t)
}
