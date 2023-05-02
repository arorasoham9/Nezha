package API

import (
	"testing"
)

func assertEquals(t *testing.T, desc string, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("%v Got: %v (%T), Want:%v (%T)", desc, got, got, want, want)
	}
}

func TestInvalidToken(t *testing.T) {
	TESTING_KEY = "TESTING"
	token := "gibberish"
	_, err := ValidateToken(token)

	if err == "" {
		t.Errorf("Validating invalid token, did not throw err")
	}
}
