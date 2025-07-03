package replicaset

import (
	"io"
	"sync"

	"github.com/mcs-unity/replica/internal/replica"
)

const ERROR_ROOT_NIL = "dir os.Root can't be a nil pointer"
const WARNING = "replication warning: log is a nil pointer errors will not be logged"

type configFile = string
type writer = func(r replica.IReplica) error
type OnError func(re replica.IReplica, err error)

type config struct {
	Url  string `json:"url,omitempty"`
	Auth string `json:"auth,omitempty"`
}

type IReplicaSet interface {
	Add(config) error
	List() []replica.IReplica
	Sync(w writer, up bool) error
}

type ReplicaSet struct {
	rw      io.Writer
	address []replica.IReplica
	lock    sync.Locker
}
