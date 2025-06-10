package decoder

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/mcs-unity/replica/internal/shared"
)

type Test struct {
	Name string
}

func TestDecode(t *testing.T) {
	r := bytes.NewBuffer([]byte{})
	p := Test{}

	go func() {
		r.Write([]byte("{\"Name\":\"TestMe\"}"))
	}()

	time.Sleep(100 * time.Millisecond)
	if err := Decode(r, &p); err != nil {
		t.Error(err)
	}

	shared.ExpectedStr(p.Name, "TestMe", t)
}

func TestEncode(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	p := Test{Name: "TestMe"}
	go func() {
		if err := Encode(b, p); err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(100 * time.Millisecond)
	buff, err := io.ReadAll(b)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Contains(buff, []byte("TestMe")) {
		t.Error("failed to find TestMe in string")
	}
}

func BenchmarkDecode(b *testing.B) {
	buff := bytes.NewBuffer([]byte{})
	p := Test{Name: "TestMe"}
	for b.Loop() {
		if err := Encode(buff, p); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	buff := bytes.NewBuffer([]byte{})
	p := Test{}

	for b.Loop() {
		go func() {
			buff.Write([]byte("{\"Name\":\"TestMe\"}"))
		}()
		time.Sleep(10 * time.Millisecond)
		if err := Decode(buff, &p); err != nil {
			b.Error(err)
		}
	}
}
