package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type PostGresRepo interface {
	Close()
}

type PostGres struct {
	DB *sql.DB
}

// initTCPConnectionPool initializes a TCP connection pool for a Cloud SQL
// instance of SQL Server.
func NewPostGresConnPool(dbUser string, dbPwd string, dbTCPHost string, dbPort string, dbName string) (*PostGres, error) {

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s sslmode=disable", dbTCPHost, dbUser, dbPwd, dbPort, dbName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("postgres", dbURI)
	if err != nil {
		return &PostGres{}, fmt.Errorf("sql.Open: %v", err)
	}

	configureConnectionPool(dbPool)

	return &PostGres{
		DB: dbPool,
	}, nil

}

func configureConnectionPool(dbPool *sql.DB) {
	// [START cloud_sql_postgres_databasesql_limit]

	// Set maximum number of connections in idle connection pool.
	dbPool.SetMaxIdleConns(5)

	// Set maximum number of open connections to the database.
	dbPool.SetMaxOpenConns(7)

	// [END cloud_sql_postgres_databasesql_limit]

	// [START cloud_sql_postgres_databasesql_lifetime]

	// Set Maximum time (in seconds) that a connection can remain open.
	dbPool.SetConnMaxLifetime(1800 * time.Second)

	// [END cloud_sql_postgres_databasesql_lifetime]
}

//
func (p *PostGres) Close() {
	p.DB.Close()
}
