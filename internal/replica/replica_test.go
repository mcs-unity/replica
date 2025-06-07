package replica

import (
	"testing"

	"github.com/mcs-unity/replica/internal/shared"
)

const ip = "180.22.31.14"
const uri = "test.com/test?param=123&param2=aa"

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
