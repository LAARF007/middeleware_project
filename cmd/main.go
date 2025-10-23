package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/events"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	//"middleware/example/internal/repositories/events"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/events", func(r chi.Router) { // route /events
		r.Get("/", events.GetEvents)          // GET /users
		r.Route("/{id}", func(r chi.Router) { // route /events/{id}
			r.Use(events.Context)       // Use Context method to get event ID
			r.Get("/", events.GetEvent) // GET /events/{id}
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

//func init() {
//	db, err := helpers.OpenDB()
//	if err != nil {
//		logrus.Fatalf("error while opening database : %s", err.Error())
//	}
//	schemes := []string{
//		`CREATE TABLE IF NOT EXISTS users (
//			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
//			name VARCHAR(255) NOT NULL
//		);`,
//	}
//	for _, scheme := range schemes {
//		if _, err := db.Exec(scheme); err != nil {
//			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
//		}
//	}
//	helpers.CloseDB(db)
//}

/*
	func init() {
		db, err := helpers.OpenDB()
		if err != nil {
			logrus.Fatalf("error while opening database : %s", err.Error())
		}
		schemes := []string{
			`CREATE TABLE IF NOT EXISTS events (
				id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
				name VARCHAR(255) NOT NULL
			);`,
		}
		for _, scheme := range schemes {
			if _, err := db.Exec(scheme); err != nil {
				logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
			}
		}
		helpers.CloseDB(db)
	}
*/
func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}

	schemes := []string{
		`CREATE TABLE IF NOT EXISTS events (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			uid VARCHAR(255),
			name VARCHAR(255) NOT NULL,
			description TEXT,
			start TIMESTAMP,
			end TIMESTAMP,
			location VARCHAR(255),
			last_update TIMESTAMP,
			agenda_ids TEXT
		);`,
	}

	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalf("Could not generate table 'events'! Error was: %s", err.Error())
		}
	}

	helpers.CloseDB(db)
}
