package replica

import (
	"errors"
	"strings"
)

func (r Replica) Address() string {
	return r.address
}

func (r Replica) State() state {
	return r.state
}

/*
provide the directory path
where there must be a replica.json
*/
func new(address string) (IReplica, error) {
	if strings.Trim(address, "") == "" {
		return nil, errors.New("empty address")
	}

	return &Replica{address: address, state: UNKNOWN}, nil
}
