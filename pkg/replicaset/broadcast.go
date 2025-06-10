package replicaset

import (
	"fmt"
)

func (r *ReplicaSet) Online() error {
	l := r.List()
	if len(l) == 0 {
		r.rw.Write([]byte("unable to replicate empty list"))
	}

	for _, re := range l {
		fmt.Println(re)
	}
	return nil
}
