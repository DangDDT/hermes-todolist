package user

import "testing"

func TestNewUserDefaultsDisplayNameToUsername(t *testing.T) {
	u, err := NewUser("dang", "")
	if err != nil {
		t.Fatalf("NewUser failed: %v", err)
	}
	if u.DisplayName != "dang" {
		t.Fatalf("expected display name fallback, got %q", u.DisplayName)
	}
}

func TestNewUserRequiresUsername(t *testing.T) {
	if _, err := NewUser("", "Dang"); err == nil {
		t.Fatal("expected username validation error")
	}
}

func TestPasswordHashAndVerify(t *testing.T) {
	u, err := NewUser("dang", "Dang")
	if err != nil {
		t.Fatal(err)
	}
	if err := u.SetPassword("secret123"); err != nil {
		t.Fatalf("SetPassword failed: %v", err)
	}
	if u.PasswordHash == "" || u.PasswordHash == "secret123" {
		t.Fatalf("expected hashed password, got %q", u.PasswordHash)
	}
	if err := u.CheckPassword("secret123"); err != nil {
		t.Fatalf("expected password to verify: %v", err)
	}
	if err := u.CheckPassword("wrong"); err == nil {
		t.Fatal("expected wrong password to fail")
	}
}
