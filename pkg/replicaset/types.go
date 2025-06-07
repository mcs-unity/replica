package replicaset

import (
	"sync"

	"github.com/mcs-unity/replica/internal/replica"
)

type configFile = string

type IReplicaSet interface {
	Add(string) error
	List() []replica.IReplica
}

type ReplicaSet struct {
	address []replica.IReplica
	lock    sync.Locker
}
