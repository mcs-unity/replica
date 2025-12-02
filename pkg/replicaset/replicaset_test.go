package replicaset

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/mcs-unity/replica/internal/shared"
	"github.com/mcs-unity/replica/pkg/replica"
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

	fn := func(r replica.IReplica) error {
		rs := &replica.RemoteState{Online: true, Timestamp: time.Now().UTC()}
		buff := shared.WriteBuffer(rs, t)
		time.Sleep(100 * time.Millisecond)
		if err := r.Online(buff); err != nil {
			return err
		}

		if r.State() != shared.UP {
			return errors.New("failed to mark replicas as online")
		}
		return nil
	}

	if err := re.Sync(fn, false); err != nil {
		t.Error(err)
	}
}

func TestAuthentication(t *testing.T) {
	r := getRoot(t)
	re, err := New(r, os.Stderr)
	if err != nil {
		t.Error(err)
	}

	fn := func(r replica.IReplica) error {
		if _, err := r.AuthKey(); err != nil {
			t.Error(err)
		}
		return nil
	}

	if err := re.Sync(fn, false); err != nil {
		t.Error(err)
	}
}
