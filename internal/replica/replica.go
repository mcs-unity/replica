package replica

import (
	"errors"
	"strings"

	"github.com/mcs-unity/replica/internal/shared"
)

func (r Replica) Address() string {
	return r.address
}

func (r Replica) State() shared.State {
	return r.state
}

/*
provide the directory path
where there must be a replica.json
*/
func New(address string) (IReplica, error) {
	if strings.Trim(address, "") == "" {
		return nil, errors.New("empty address")
	}

	return &Replica{address: address, state: shared.UNKNOWN}, nil
}
