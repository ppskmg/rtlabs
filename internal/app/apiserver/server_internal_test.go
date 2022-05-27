package apiserver

import (
	"rtlabs/internal/app/store/postgresstore"
	"testing"
)

var (
	databaseURL = "host=localhost dbname=rtlabs_test sslmode=disable"
)

type testClient struct {
	testServer  *server
	testClients Client
}

// testServer tc, teardownPostgres, teardownRedis, teardownRedisBlackList, teardownRedisRefreshBlackList
func testServer(t *testing.T) (tc *testClient, teardownPostgres func(...string)) {

	db, teardownPostgres := postgresstore.TestDB(t, databaseURL)

	sc := Client{
		Postgres: postgresstore.New(db),
	}
	tc = &testClient{
		testServer:  newServer(sc),
		testClients: sc,
	}

	return tc, teardownPostgres

}
