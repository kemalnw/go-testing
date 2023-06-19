package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserNameByID(t *testing.T) {
	// Open a database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	repo := NewUserRepository(db)

	// Test case 1: Success case
	mock.ExpectQuery("SELECT username FROM users WHERE id = ?").WithArgs(1).WillReturnRows(mockRows("Kemal"))
	name, err := repo.GetUserNameByID(1)

	assert.Nil(t, err)
	assert.Equal(t, "Kemal", name)

	// Test case 2: Error case
	mock.ExpectQuery("SELECT username FROM users WHERE id = ?").WithArgs(1).WillReturnError(errors.New("fail"))
	name, err = repo.GetUserNameByID(1)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Empty(t, name)
}

func mockRows(name string) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"username"}).AddRow(name)
}
