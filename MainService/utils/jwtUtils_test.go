package utils

import (
	"server/MainService/config"
	"testing"

	"github.com/stretchr/testify/require"
)

type TConfig struct {
}

func (t TConfig) GetConfig(key config.ConfigKey) interface{} {
	if key == config.HMAC_KEY {
		return "Tam"
	}
	return ""
}

func TestDecodeToken(t *testing.T) {
	j := JwtUtils{c: TConfig{}}
	want := InfoInJwt{UserId: 202}
	token, err := j.GenerateToken(want)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	got, err := j.DecodeToken(token)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}

	if !equalCredentials(*got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestGenerateToken(t *testing.T) {
	j := JwtUtils{c: TConfig{}}
	want := InfoInJwt{UserId: 1}
	token, err := j.GenerateToken(want)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	got, err := j.DecodeToken(token)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}

	if !equalCredentials(*got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}

func equalCredentials(got, want InfoInJwt) bool {
	return got.UserId == want.UserId
}

func TestVerifyToken(t *testing.T) {
	j := JwtUtils{c: TConfig{}}
	want := InfoInJwt{UserId: 1}
	token, err := j.GenerateToken(want)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
	err = j.VerifyToken(token)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}

	require.NoError(t, err)
}
