package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"middleware/example/internal/controllers"
	"middleware/example/internal/helpers"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	// Routes Alerts
	r.Route("/alerts", func(r chi.Router) {
		r.Get("/", controllers.GetAllAlerts) // GET /alerts
		r.Post("/", controllers.CreateAlert)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(controllers.ContextIDs("alertId")) // Middleware pour récupérer alertId
			r.Get("/", controllers.GetAlertByID)
			r.Delete("/", controllers.DeleteAlert)
			r.Put("/", controllers.UpdateAlert)
		})
	})

	// Routes Agendas
	r.Route("/agendas", func(r chi.Router) {
		r.Get("/", controllers.GetAllAgendas) // GET /agendas
		r.Post("/", controllers.CreateAgenda)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(controllers.ContextIDs("agendaId")) // Middleware pour récupérer agendaId
			r.Get("/", controllers.GetAgendaByID)
			r.Delete("/", controllers.DeleteAgenda)
			r.Put("/", controllers.UpdateAgenda)

		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}

	// Création des tables si elles n'existent pas
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS agendas (
			id VARCHAR(3) PRIMARY KEY NOT NULL UNIQUE,
			ucaId INTEGER NOT NULL,
			name VARCHAR(255) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS alerts (
			id VARCHAR(3) PRIMARY KEY NOT NULL UNIQUE,
			email VARCHAR(255) NOT NULL,
			agendaId VARCHAR(3) NOT NULL,
			FOREIGN KEY(agendaId) REFERENCES agendas(id)
		);`,
	}

	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}

	query := `INSERT INTO agendas (id, ucaId, name) VALUES (?, ?, ?)`

	id := uuid.Must(uuid.NewV4())
	_, err = db.Exec(query,
		id.String(),
		"56529",
		"M1 - Tutorat L2",
	)

	if err != nil {
		logrus.Warnf("Could not insert agenda: %v", err)
	}

	fmt.Println("Agenda inséré avec succès dans SQLite !")

	//Alerts

	/*query := `INSERT INTO alerts (id, email, agendaId) VALUES (?, ?, ?)`

	id := uuid.Must(uuid.NewV4())
	_, err = db.Exec(query,
		id.String(),
		"56529",
		"M1 - Tutorat L2",
	)

	if err != nil {
		logrus.Warnf("Could not insert alert: %v", err)
	}

	fmt.Println("Alerts inséré avec succès dans SQLite !")*/

	/*if count == 0 {
		// IDs simples de 3 caractères
		agendaIDs := []string{"A01", "B02", "C03", "D04"}
		alertIDs := []string{"X01", "Y02", "Z03", "W04"}

		// Insertion des agendas et des alerts
		for i := 0; i < 4; i++ {
			_, err := db.Exec(`INSERT INTO agendas (id, ucaId, name) VALUES (?, ?, ?)`,
				agendaIDs[i], i+1, fmt.Sprintf("Agenda %d", i+1))
			if err != nil {
				logrus.Warnf("Could not insert agenda: %v", err)
			}

			_, err = db.Exec(`INSERT INTO alerts (id, email, agendaId) VALUES (?, ?, ?)`,
				alertIDs[i], fmt.Sprintf("user%d@example.com", i+1), agendaIDs[i])
			if err != nil {
				logrus.Warnf("Could not insert alert: %v", err)
			}
		}

		logrus.Info("Initial data inserted into agendas and alerts")
	} else {
		logrus.Info("Agendas table already populated, skipping initial data insert")
	}*/

	helpers.CloseDB(db)
}
