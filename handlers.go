package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

// request header: -H "Authorization: Basic d57be6b2402fe1248ca912bccd061653f743cf21b4c382f6b" -H "freckle-platform: android" -H "freckle-sdk-version: 1.1" -H "freckle-app-id: com.freckleiot.freckle" -H "freckle-idfa: tester-12345-12345" -H "freckle-user-lang: en" -H "Content-Type: application/json"

func Index(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	action := r.Form.Get("action")
	fmt.Fprintln(w, "Welcome!", action)

	switch action {
	case "ping", "Ping":
		fmt.Fprintln(w, "We've gone a ", action, " event")
	case "list", "List":

		beacons := Beacons{}

		app_id := r.Header.Get("freckle-app-id")
		lat := r.Form.Get("lat")
		lng := r.Form.Get("lng")
		fmt.Fprintln(w, "app-id: ", app_id,
			"lat:", lat,
			"lng:", lng)

		//selectAll(&beacons)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(beacons); err != nil {
			log.Fatal(err)
			panic(err)
		}

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

//
// /Status returns HTTP 200 to signal "I am alive and well!"
//
func Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if hostname, err := os.Hostname(); err == nil {
		fmt.Fprintln(w, "Success! I am", hostname, ". Server time:", time.Now())
	} else {
		fmt.Fprintln(w, err.Error)
	}
}

//
// List
//
func List(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	option := r.Form.Get("option")
	switch option {

	case "latlng":

		fmt.Println("------------------------------------------")

		// Parse lat & log without knowing whether the exist or not
		latStr := r.Form.Get("lat")
		lngStr := r.Form.Get("lng")

		fmt.Fprintln(w, latStr, lngStr)
		beacons := selectProximity_P_Beacons(43.6536106, -79.3800603)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(beacons); err != nil {
			panic(err)
		}

	// list all beacons straight from DB
	case "from-db":
		beacons := select_all()
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(beacons); err != nil {
			panic(err)
		}

	// list all beacons from local memory
	case "echo":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

	// list all beacons from local memory
	case "text":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v", CachedBeacons)

	// list all beacons from local memory
	case "from-cache":
		fallthrough

	// list all beacons from local memory
	default:
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(CachedBeacons); err != nil {
			panic(err)
		}

	}

}
