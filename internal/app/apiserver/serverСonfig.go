package apiserver

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port             string
	PortTLS     string
	SSLCrt string
	SSLKey string
	LogLevel         string
	DatabaseURL      string
	DatabaseURLMongo string
	DatabaseName     string
	DatabaseURLTest  string
	DatabaseNameTest string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	// TODO: err
	port, _ := os.LookupEnv("PORT")
	portTls, _ := os.LookupEnv("PORT_TLS")
	sslCrt, _ := os.LookupEnv("SSL_CRT")
	sslKey, _ := os.LookupEnv("SSL_KEY")
	dbPG, _ := os.LookupEnv("DB_PG")
	dbMongo, _ := os.LookupEnv("DB_MONGO")
	dbName, _ := os.LookupEnv("DBNAME")
	dbTest, _ := os.LookupEnv("DB_TEST")
	dbNameTest, _ := os.LookupEnv("DB_NAME_TEST")

	return &Config{
		Port:        port,
		PortTLS:     portTls,
		SSLCrt:      sslCrt,
		SSLKey:      sslKey,
		LogLevel:    "debug",
		DatabaseURL: dbPG,
		//DatabaseURL:      "host=localhost dbname=rtlabs sslmode=disable",
		DatabaseURLMongo: dbMongo,
		DatabaseName:     dbName,
		DatabaseURLTest:  dbTest,
		DatabaseNameTest: dbNameTest,
	}
}
