package replicaset

import (
	"os"
	"testing"
)

func getRoot(t *testing.T) *os.Root {
	r, err := os.OpenRoot("./")
	if err != nil {
		t.Error(err)
	}
	return r
}

func TestLoadReplicaSet(t *testing.T) {
	r := getRoot(t)
	re, err := New(r, os.Stderr)
	if err != nil {
		t.Error(err)
	}

	if len(re.List()) < 2 {
		t.Error("expected two hosts")
	}
}

func TestNilWriter(t *testing.T) {
	r := getRoot(t)
	_, err := New(r, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestBadReplica(t *testing.T) {
	_, err := New(nil, os.Stderr)
	if err == nil {
		t.Error("failed to capture nil pointer")
	}
}

func TestOnline(t *testing.T) {
	r := getRoot(t)
	re, err := New(r, os.Stderr)
	if err != nil {
		t.Error(err)
	}

	if err := re.Online(); err != nil {
		t.Error(err)
	}
}
