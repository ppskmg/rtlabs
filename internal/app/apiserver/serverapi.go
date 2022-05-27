package apiserver

import (
	"database/sql"
	"fmt"
	"net/http"
	"rtlabs/internal/app/store/postgresstore"

	_ "github.com/lib/pq"
)

type Client struct {
	Postgres *postgresstore.Store
}

func Start(config *Config, tls bool) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	postgresStore := postgresstore.New(db)

	if err != nil {
		return err
	}

	clients := Client{
		Postgres: postgresStore,
	}
	srv := newServer(clients)
	if tls {
		fmt.Println("ListenAndServeTLS: ", config.PortTLS)
		return http.ListenAndServeTLS(config.PortTLS, config.SSLCrt, config.SSLKey, srv)
		//return http.ListenAndServeTLS("localhost:5001", config.SSLCrt, config.SSLKey, srv)
	}

	return http.ListenAndServe(config.Port, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
