package replica

import (
	"io"
	"net/url"
	"time"

	"github.com/mcs-unity/replica/internal/shared"
)

type IReplica interface {
	Address() *url.URL
	AuthKey() (string, error)
	State() shared.State
	Report(s shared.State)
	Online(rw io.Reader) error
}

type Replica struct {
	addr  *url.URL
	auth  string
	state shared.State
}

type RemoteState struct {
	Online    bool
	Timestamp time.Time
}
