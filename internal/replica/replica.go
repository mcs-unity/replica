package replica

import (
	"errors"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/mcs-unity/replica/internal/decoder"
	"github.com/mcs-unity/replica/internal/shared"
	"github.com/mcs-unity/replica/pkg/remotetypes"
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
returns latest error message that was reported
*/
func (r Replica) Error() ErrorMessage {
	return r.err
}

/*
assigns an error to the errorMessage
*/
func (r *Replica) SetError(err error) {
	if err == nil {
		return
	}
	r.err = ErrorMessage{time.Now().UTC(), err}
}

/*
online will read a buffer expecting a json string
the expected payload is a RemoteState
*/
func (r *Replica) Online(re io.Reader) error {
	rs := &remotetypes.RemoteState{}
	if err := decoder.Decode(re, rs); err != nil {
		r.state = shared.UNKNOWN
		return err
	}

	if rs.Online {
		r.state = shared.UP
	}
	return nil
}

/*
returns a authentication key that can be used
in a HTTP authorization header
*/
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
