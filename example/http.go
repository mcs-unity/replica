package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/mcs-unity/replica/pkg/remotetypes"
)

func online(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
	}

	resp, err := json.Marshal(remotetypes.RemoteState{Online: false, Timestamp: time.Now().UTC()})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("failed to process payload"))
	}

	w.WriteHeader(200)
	w.Write(resp)
}

func startHttp() {
	log.Println("Starting http server")
	handler := http.NewServeMux()
	handler.HandleFunc("/online", online)
	if err := http.ListenAndServe("localhost:3000", handler); err != nil {
		panic(err)
	}
}
