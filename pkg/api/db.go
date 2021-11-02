package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// go mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type mysqlConfig struct {
	// MySQL user
	User string

	// MySQL password
	Pass string

	// MySQL database
	DB string

	// MySQL host
	Host string

	// MySQL port
	Port string
}

var mysqlConf mysqlConfig

// SetupMySQLConfig updates mysqlConf using environment variables.
func SetupMySQLConfig() {
	// Set user
	mysqlConf.User = "brucewayne"

	// Set password
	mysqlConf.Pass = "batman"

	// Set database
	mysqlConf.DB = "superhero"

	// Set host
	if mysqlConf.Host = os.Getenv("MYSQL_SERVICE_HOST"); mysqlConf.Host == "" {
		log.Fatal("MYSQL_SERVICE_HOST environment variable is missing")
	}

	// Set port
	if mysqlConf.Port = os.Getenv("MYSQL_SERVICE_PORT"); mysqlConf.Port == "" {
		log.Fatal("MYSQL_SERVICE_PORT environment variable is missing")
	}
}

// dataSourceName returns a connection string suitable for sql.Open.
// See https://github.com/go-sql-driver/mysql#dsn-data-source-name
func (c mysqlConfig) dataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Pass, c.Host, c.Port, c.DB)
}

// Database creates a new database connection.
func Database() (*sql.DB, error) {
	db, err := sql.Open("mysql", mysqlConf.dataSourceName())
	if err != nil {
		return nil, fmt.Errorf("mysql: could not get a connection: %v", err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("mysql: could not establish a good connection: %v", err)
	}

	return db, nil
}
