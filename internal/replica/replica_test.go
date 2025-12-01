package replica

import (
	"testing"
	"time"

	"github.com/mcs-unity/replica/internal/shared"
	"github.com/mcs-unity/replica/pkg/remotetypes"
)

const ip = "http://192.168.1.1"
const uri = "https://test.com/test?param=123&param2=aa"

func new(address string, t *testing.T) IReplica {
	r, err := New(address, "")
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
	_, err := New("", "")
	if err == nil {
		t.Error("failed to capture empty string")
	}

	_, err = New("192.168.1", "")
	if err == nil {
		t.Error("failed to capture bad ip address")
	}

	_, err = New("domain.com?/asd", "")
	if err == nil {
		t.Error("failed to capture bad url")
	}
}

func TestOnline(t *testing.T) {
	r := new(ip, t)
	shared.IsNil(r, "replica", t)
	rs := &remotetypes.RemoteState{Online: true, Timestamp: time.Now().UTC()}
	buff := shared.WriteBuffer(rs, t)
	time.Sleep(10 * time.Millisecond)
	if err := r.Online(buff); err != nil {
		t.Error(err)
	}

	if r.State() != shared.UP {
		t.Error("failed to process state")
	}
}

func TestGetAuthKey(t *testing.T) {
	key := "samplekey"
	r, err := New(ip, key)
	if err != nil {
		t.Error(err)
	}

	auth, err := r.AuthKey()
	if err != nil {
		t.Error(err)
	}

	shared.ExpectedStr(auth, key, t)
}

func TestEmptyAuthKey(t *testing.T) {
	r, err := New(ip, "")
	if err != nil {
		t.Error(err)
	}

	if _, err := r.AuthKey(); err == nil {
		t.Error("empty auth key error expected")
	}
}
