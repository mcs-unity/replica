package replica

import (
	"github.com/mcs-unity/replica/internal/shared"
)

type IReplica interface {
	Address() string
	State() shared.State
}

type Replica struct {
	address string
	state   shared.State
}
