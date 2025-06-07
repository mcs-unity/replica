package replica

import (
	"net/url"

	"github.com/mcs-unity/replica/internal/shared"
)

type IReplica interface {
	Address() *url.URL
	State() shared.State
}

type Replica struct {
	address *url.URL
	state   shared.State
}
