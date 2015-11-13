package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	action := r.Form.Get("action")
	fmt.Fprintln(w, "Welcome!", action)
	switch action {
	case "ping", "Ping":
		fmt.Fprintln(w, "We've gone a ", action, " event")
	case "list", "List":
		fmt.Fprintln(w, "We've gone a", action, "event!")
	case "enter", "Enter":
		fmt.Fprintln(w, "We've gone a ", action, " event")
	case "exit", "Exit":
		fmt.Fprintln(w, "We've gone a ", action, " event")
	case "notified", "Notified":
		fmt.Fprintln(w, "We've gone a ", action, " event")
	default:
		fmt.Fprintln(w, "Unknown action")
	}
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

// /Status returns HTTP 200 to signal "I am alive and well!"
func Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if hostname, err := os.Hostname(); err == nil {
		fmt.Fprintln(w, "Success! I am", hostname, ". Server time:", time.Now())
	} else {
		fmt.Fprintln(w, err.Error)
	}
}
