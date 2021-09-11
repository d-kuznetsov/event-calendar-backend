package service

import (
	"testing"
)

func TestToken(t *testing.T) {
	initPayload := "1234567"
	tokenStr, err := generateToken(initPayload)
	if err != nil {
		t.Error("Token generating error:", err)
	}
	payload, err := parseToken(tokenStr)
	if err != nil {
		t.Error("Token parsing error:", err)
	}
	if payload != initPayload {
		t.Error("Wrong payload, was expecting " + initPayload + " but got " + payload)
	}
}
