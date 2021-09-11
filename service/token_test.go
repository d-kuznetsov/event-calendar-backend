package service

import (
	"testing"
)

func TestToken(t *testing.T) {
	initEmail := "email@test"
	tokenStr, err := generateToken(initEmail)
	if err != nil {
		t.Error("Token generating error:", err)
	}
	email, err := parseToken(tokenStr)
	if err != nil {
		t.Error("Token parsing error:", err)
	}
	if email != initEmail {
		t.Error("Wrong email, was expecting " + initEmail + " but got " + email)
	}
}
