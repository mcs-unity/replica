package replicaset

import (
	"os"
	"testing"
)

func TestLoadReplicaSet(t *testing.T) {
	r, err := os.OpenRoot("./")
	if err != nil {
		t.Error(err)
	}

	_, err = New(r)
	if err != nil {
		t.Error(err)
	}
}
