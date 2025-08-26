package replicaset

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/mcs-unity/replica/internal/replica"
)

/*
creates a new replica and adds it into the replicaset
*/
func (r *ReplicaSet) Add(c config) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	replica, err := replica.New(c.Url, c.Auth)
	if err != nil {
		return err
	}

	r.address = append(r.address, replica)
	return nil
}

/*
returns a list of replicas registered
*/
func (r *ReplicaSet) List() []replica.IReplica {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.address
}

func (r ReplicaSet) Log(err error) {
	if r.rw == nil {
		return
	}

	if _, err := r.rw.Write([]byte(err.Error() + "\n")); err != nil {
		fmt.Println(err)
	}
}

/*
read the file and decode the json value into the replica set
*/
func (re *ReplicaSet) processFile(r io.Reader) error {
	read := json.NewDecoder(r)
	arr := make([]config, 0)
	if err := read.Decode(&arr); err != nil {
		return err
	}

	for _, l := range arr {
		if err := re.Add(l); err != nil {
			re.Log(err)
		}
	}
	return nil
}

/*
provide the directory path
where there must be a replica.json
*/
func New(dir *os.Root, log io.Writer) (IReplicaSet, error) {
	if dir == nil {
		return nil, errors.New(ERROR_ROOT_NIL)
	}

	if log == nil {
		fmt.Println(WARNING)
	}

	file, err := dir.OpenFile("replica.json", os.O_RDONLY, 0o777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	set := &ReplicaSet{rw: log, address: []replica.IReplica{}, lock: &sync.Mutex{}}
	if err := set.processFile(file); err != nil {
		return nil, err
	}
	return set, nil
}
