package replica

import "sync"

type configFile = string
type state uint8

const (
	DOWN    state = iota
	UP      state = iota
	UNKNOWN state = iota
)

type IReplicaSet interface {
	Add(string) error
	List() []IReplica
}

type IReplica interface {
	Address() string
	State() state
}

type Replica struct {
	address string
	state
}

type Replicas struct {
	address []IReplica
	lock    sync.Locker
}
