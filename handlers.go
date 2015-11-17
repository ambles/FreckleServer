package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
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

		beacons := Beacons{}

		//selectAll(&beacons)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(beacons); err != nil {
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

func selectAll(beacons *Beacons) {
	// TODO : user and dbname should be in configuration
	db, err := sql.Open("postgres", "user=postgres dbname=freckle_proximity_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// TODO : 'beacons' should be a variable, and do we need '*', be specific!
	rows, err := db.Query("SELECT * FROM beacons")
	if err != nil {
		panic(err.Error())
	}

	if err = db.Close(); err != nil {
		panic(err.Error())
	}

	var (
		beaconid         string
		date             time.Time
		uuid             string
		major            int
		minor            int
		nickname         []byte
		current_campaign []byte
		inredis          bool
		active           bool
		lat              float64
		long             float64
		location         []byte
		tags             []byte
		attributes       []byte
		class            []byte
		creation_date    time.Time
		geom             []byte
	)

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&beaconid,         // 1
			&date,             // 2
			&uuid,             // 3
			&major,            // 4
			&minor,            // 5
			&nickname,         // 6
			&current_campaign, // 7
			&inredis,          // 8
			&active,           // 9
			&lat,              // 10
			&long,             // 11
			&location,         // 12
			&tags,             // 13
			&attributes,       // 14
			&class,            // 15
			&creation_date,    // 16
			&geom,             // 17
		)
		if err != nil {
			log.Fatal(err)
		}

		// TODO : IO here slows things down
		log.Println(beaconid)

		// TODO : the following append is VERY memory inefficient
		*beacons = append(*beacons, Beacon{ID: beaconid, CreationDate: creation_date})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
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
