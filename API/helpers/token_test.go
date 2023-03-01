package helpers

import (
	"os"
	"testing"
)

func assertEquals(t *testing.T, desc string, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("%v Got: %v (%T), Want:%v (%T)", desc, got, got, want, want)
	}
}

func TestInvalidToken(t *testing.T) {
	os.Setenv("SECRET_KEY", "SECRET_KEY")
	token := "gibberish"
	os.Setenv("SECRET_KEY", "SECRET_KEY")
	_, err := ValidateToken(token)

	if err == "" {
		t.Errorf("Validating invalid token, did not throw err")
	}
}
