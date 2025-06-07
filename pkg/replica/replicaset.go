package replica

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

func (r *Replicas) Add(address string) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	replica, err := new(address)
	if err != nil {
		return err
	}

	r.address = append(r.address, replica)

	return nil
}

func (r Replicas) List() []IReplica {
	return r.address
}

/*
read the file and decode the json value into the replica set
*/
func processFile(r io.Reader, set IReplicaSet) error {
	read := json.NewDecoder(r)

	if err := read.Decode(&set); err != nil {
		return err
	}

	return nil
}

/*
provide the directory path
where there must be a replica.json
*/
func New(dir os.Root) (IReplicaSet, error) {
	file, err := dir.OpenFile("replica.json", os.O_RDONLY, 0o777)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	set := &Replicas{address: []IReplica{}, lock: &sync.Mutex{}}
	if err := processFile(file, set); err != nil {
		return nil, err
	}

	return nil, nil
}
