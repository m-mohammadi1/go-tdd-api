package postgres

import (
	"testing"

	"github.com/m-mohammadi1/go-json-api/pkg/domain"
)

func TestCreateUser(t *testing.T) {
	pStore := NewPostgresStore(testDB)

	oldPassword := "password"
	user := &domain.User{
		Email:    "test@test.com",
		Password: oldPassword,
		Name:     "Mohammad Mohammadi",
	}

	createdUser, err := pStore.CreateUser(user)

	if err != nil {
		t.Fatal(err)
	}

	if createdUser.ID == 0 {
		t.Errorf("user ID equals 0")
	}

	if user.Name != createdUser.Name {
		t.Errorf("Excpected %q; got %q", user.Name, createdUser.Name)
	}

	if user.Password == createdUser.Name {
		t.Errorf("password not hashed")
	}
}
