package shared

import (
	"bytes"
	"encoding/json"
	"io"
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
		t.FailNow()
	}
}

func WriteBuffer(payload any, t *testing.T) io.Reader {
	buff := bytes.NewBuffer([]byte{})
	go func() {
		b, err := json.Marshal(payload)
		if err != nil {
			t.Error(err)
		}
		if _, err := buff.Write(b); err != nil {
			t.Error(err)
		}
	}()
	return buff
}
