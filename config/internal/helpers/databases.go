package helpers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

// OpenDB ouvre la connexion SQLite pour ton projet
func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "agendas_alerts.db")
	if err != nil {
		logrus.Errorf("Impossible d'ouvrir la base : %s", err.Error())
		return nil, err
	}

	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		logrus.Errorf("Erreur de connexion à la base : %s", err.Error())
		return nil, err
	}

	return db, nil
}

// CloseDB ferme la connexion à la base de données
func CloseDB(db *sql.DB) {
	if db == nil {
		return
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("Erreur lors de la fermeture de la base : %s", err.Error())
	}
}
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	body, _ := json.Marshal(data)
	_, _ = w.Write(body)
}
