package shared

import (
	"testing"
)

func ExpectedStr(a string, b string, t *testing.T) {
	if a != b {
		t.Errorf("expected %s got %s", a, b)
	}
}

func ExpectedInt(a int, b int, t *testing.T) {
	if a != b {
		t.Errorf("expected %d got %d", a, b)
	}
}

func IsNil(a any, name string, t *testing.T) {
	if a == nil {
		t.Errorf("%s is nil", name)
	}
}
