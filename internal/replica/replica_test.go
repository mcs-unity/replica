package replica

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/mcs-unity/replica/internal/shared"
)

const ip = "http://192.168.1.1"
const uri = "https://test.com/test?param=123&param2=aa"

func new(address string, t *testing.T) IReplica {
	r, err := New(address)
	if err != nil {
		t.Error(err)
	}

	shared.IsNil(r, "replica", t)
	return r
}

func TestNewIP(t *testing.T) {
	new(ip, t)
}

func TestNewUrl(t *testing.T) {
	new(uri, t)
}

func TestGetAddress(t *testing.T) {
	r := new(ip, t)
	shared.IsNil(r, "replica", t)

	shared.ExpectedStr(r.Address().String(), ip, t)
}

func TestGetInitialState(t *testing.T) {
	r := new(ip, t)
	shared.IsNil(r, "replica", t)
	shared.ExpectedInt(int(r.State()), int(shared.UNKNOWN), t)
}

func TestBadIp(t *testing.T) {
	_, err := New("")
	if err == nil {
		t.Error("failed to capture empty string")
	}

	_, err = New("192.168.1")
	if err == nil {
		t.Error("failed to capture bad ip address")
	}

	_, err = New("domain.com?/asd")
	if err == nil {
		t.Error("failed to capture bad url")
	}
}

func TestOnline(t *testing.T) {
	r := new(ip, t)
	shared.IsNil(r, "replica", t)

	rs := &RemoteState{Online: true, timestamp: time.Now().UTC()}
	buff := bytes.NewBuffer([]byte{})
	go func() {
		b, err := json.Marshal(rs)
		if err != nil {
			t.Error(err)
		}
		if _, err := buff.Write(b); err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(10 * time.Millisecond)
	if err := r.Online(buff); err != nil {
		t.Error(err)
	}

	if r.State() != shared.UP {
		t.Error("failed to process state")
	}
}
