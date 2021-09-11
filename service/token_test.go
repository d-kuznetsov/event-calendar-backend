package service

import (
	"testing"
)

func TestToken(t *testing.T) {
	initPayload := "email@test"
	tokenStr, err := generateToken(initPayload)
	if err != nil {
		t.Error("Token generating error:", err)
	}
	payload, err := parseToken(tokenStr)
	if err != nil {
		t.Error("Token parsing error:", err)
	}
	if payload != initPayload {
		t.Error("Wrong email, was expecting " + initPayload + " but got " + payload)
	}
}
