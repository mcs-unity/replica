package replica

import (
	"errors"
	"net/url"
	"strings"

	"github.com/mcs-unity/replica/internal/shared"
)

/*
returns an ipv4 address or url
*/
func (r Replica) Address() *url.URL {
	return r.address
}

/*
returns the state or the replica
*/
func (r Replica) State() shared.State {
	return r.state
}

/*
provide ip or valid url to create a new replica
example:
ip: 192.168.32.33/path/:id?param1=pm&param2=pm2 or
url: mydomain.com/path/:id?param1=pm&param2=pm2
*/
func New(address string) (IReplica, error) {
	if strings.Trim(address, "") == "" {
		return nil, errors.New("empty address")
	}

	url, err := url.Parse(address)
	if err != nil {
		return nil, err
	}

	return &Replica{address: url, state: shared.UNKNOWN}, nil
}
