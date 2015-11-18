package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

//
//
//
func selectProximity_P_Beacons(lat float32, lng float32) Beacons {

	var beacons Beacons

	db, err := sql.Open("postgres", "user=postgres dbname=freckle_proximity_db sslmode=disable")
	if err != nil {
		if db != nil {
			defer db.Close() // ignoring err as we are in error state already
		}
		log.Fatal(err)
	}

	// TODO : 'beacons' should be a variable, and do we need '*', be specific!
	//rows, err := db.Query("SELECT b.uuid, b.major, b.minor, b.beaconid AS id, ROUND((b.geom <-> ST_MakePoint(?, -79.3800603))::numeric(10, 4), 0)::float AS distance_in_meters FROM beacons b WHERE b.lat IS NOT NULL AND b.lng IS NOT NULL AND b.uuid NOT LIKE 'unit%' AND b.geom <-> ST_MakePoint(43.6536106, -79.3800603) < 10000 ORDER BY distance_in_meters ASC LIMIT 20;", lat)
	rows, err := db.Query("SELECT b.uuid, b.major, b.minor, b.beaconid AS id, ROUND((b.geom <-> ST_MakePoint(43.6536106, -79.3800603))::numeric(10, 4), 0)::float AS distance_in_meters FROM beacons b WHERE b.lat IS NOT NULL AND b.lng IS NOT NULL AND b.uuid NOT LIKE 'unit%' AND b.geom <-> ST_MakePoint(43.6536106, -79.3800603) < 10000 ORDER BY distance_in_meters ASC LIMIT 20;")
	if err != nil {
		if db != nil {
			defer db.Close() // ignoring err as we are in error state already
		}
		panic(err.Error())
	}

	if err = db.Close(); err != nil {
		panic(err.Error())
	}

	var (
		beaconid string
		uuid     string
		major    int
		minor    int
		distance int
	)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&uuid,     // 2
			&major,    // 3
			&minor,    // 4
			&beaconid, // 1
			&distance, // 5
		)
		if err != nil {
			log.Fatal(err)
		}

		// TODO : the following append is VERY memory inefficient
		beacons = append(beacons, Beacon{ID: beaconid})
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return beacons
}

//
//
//
func select_all() Beacons {

	var beacons Beacons

	// TODO : user and dbname should be in configuration
	db, err := sql.Open("postgres", "user=postgres dbname=freckle_proximity_db sslmode=disable")
	if err != nil {
		if db != nil {
			defer db.Close() // ignoring err as we are in error state already
		}
		log.Fatal(err)
	}

	// TODO : 'beacons' should be a variable, and do we need '*', be specific!
	rows, err := db.Query("SELECT * FROM beacons")
	if err != nil {
		if db != nil {
			defer db.Close() // ignoring err as we are in error state already
		}
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

		// TODO : the following append is VERY memory inefficient
		beacons = append(beacons, Beacon{ID: beaconid})
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return beacons
}
