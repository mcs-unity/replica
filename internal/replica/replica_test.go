package replica

import (
	"testing"

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
