package common

import "testing"

func TestPasswordHash(t *testing.T) {
	password := "password"

	hashed, err := PasswordHash(password)

	if err != nil {
		t.Fatal(err)
	}

	if len(hashed) == 0 {
		t.Errorf("want hash; got %q", hashed)
	}

	if hashed == password {
		t.Errorf("password was not hashed; got %q", hashed)
	}
}

func TestCheckPassword(t *testing.T) {
	password := "password"

	hashed, err := PasswordHash(password)
	if err != nil {
		t.Fatal(err)
	}

	err = CheckPassword(password, hashed)

	if err != nil {
		t.Errorf("password verification failed; got %q", err)
	}
}

func TestPasswordHashError(t *testing.T) {
	longPassword := make([]byte, 73)
	_, err := PasswordHash(string(longPassword))

	if err == nil {
		t.Fatal("expected an error; got nil")
	}
}
