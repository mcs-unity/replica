package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mcs-unity/replica/pkg/replica"
	"github.com/mcs-unity/replica/pkg/replicaset"
)

func isOnline(r replica.IReplica) error {
	resp, err := http.Get("http://localhost:3000/online")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("unable to process request, %d:%s", resp.StatusCode, body)
	}

	if err := r.Online(resp.Body); err != nil {
		return err
	}

	return nil
}

func start() replicaset.IReplicaSet {
	r, err := os.OpenRoot("./")
	if err != nil {
		panic(err)
	}

	rep, err := replicaset.New(r, os.Stdout)
	if err != nil {
		panic(err)
	}
	return rep
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	go startHttp()
	rep := start()
	log.Println("sleep for 500ms to give time for the http server to be ready")
	time.Sleep(500 * time.Millisecond)
	log.Println("begin synchronization")
	if err := rep.Sync(isOnline, false); err != nil {
		log.Println(err)
	}

	log.Println("Sync completed presenting payloads")
	for _, r := range rep.List() {
		fmt.Printf("url: %s state: %d\n", r.Address(), r.State())
	}
}
