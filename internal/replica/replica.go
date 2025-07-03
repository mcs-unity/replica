package replica

import (
	"errors"
	"io"
	"net/url"
	"strings"

	"github.com/mcs-unity/replica/internal/decoder"
	"github.com/mcs-unity/replica/internal/shared"
)

/*
returns an registered url
*/
func (r Replica) Address() *url.URL {
	return r.addr
}

/*
returns the state or the replica
*/
func (r Replica) State() shared.State {
	return r.state
}

/*
change state
*/
func (r *Replica) Report(s shared.State) {
	r.state = s
}

/*
online will read a buffer expecting a json string
the expected payload is a RemoteState
*/
func (r *Replica) Online(re io.Reader) error {
	rs := &RemoteState{}
	if err := decoder.Decode(re, rs); err != nil {
		r.state = shared.UNKNOWN
		return err
	}

	if rs.Online {
		r.state = shared.UP
	}
	return nil
}

func (r Replica) AuthKey() (string, error) {
	if strings.Trim(r.auth, "") == "" {
		return "", errors.New("empty auth key")
	}
	return r.auth, nil
}

/*
provide ip or valid url to create a new replica
example:
ip: 192.168.32.33/path/:id?param1=pm&param2=pm2 or
url: mydomain.com/path/:id?param1=pm&param2=pm2
*/
func New(addr string, authKey string) (IReplica, error) {
	if strings.Trim(addr, "") == "" {
		return nil, errors.New("empty address")
	}

	uri, err := url.ParseRequestURI(addr)
	if err != nil {
		return nil, err
	}

	return &Replica{addr: uri, auth: authKey, state: shared.UNKNOWN}, nil
}
