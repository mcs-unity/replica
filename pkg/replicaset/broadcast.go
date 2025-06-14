package replicaset

import (
	"errors"

	"github.com/mcs-unity/replica/internal/replica"
	"github.com/mcs-unity/replica/internal/shared"
)

/*
ensures that the replica list is populated and if not it
will return an error
*/
func (r ReplicaSet) isEmpty() ([]replica.IReplica, error) {
	l := r.List()
	if l == nil {
		return nil, errors.New("replica list is nil")
	}

	if len(l) == 0 {
		return nil, errors.New("unable to replicate empty list")
	}
	return l, nil
}

/*
synchronize data between system and replicas
to only send to replicas that are deemed online
set (up) to true
*/
func (r *ReplicaSet) Sync(w writer, up bool) error {
	if w == nil {
		return errors.New("writer is a nil pointer")
	}

	l, err := r.isEmpty()
	if err != nil {
		return err
	}

	if write(l, w, r.OnError, up) {
		return errors.New("failures reported verify replica states")
	}

	return nil
}

/*
register error and report the error state to the replica
*/
func (r *ReplicaSet) OnError(re replica.IReplica, err error) {
	r.Log(err)
	re.Report(shared.ERROR)
}

/*
writes to all listed replicas in the slice. to only send
to replicas that are deemed online set up to true
*/
func write(l []replica.IReplica, w writer, fn OnError, up bool) bool {
	failed := false
	for _, re := range l {
		if up && re.State() < shared.ERROR {
			continue
		}

		if err := w(re); err != nil {
			failed = true
			fn(re, err)
		}
	}
	return failed
}
