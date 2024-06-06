package Database

import (
	"database/sql"
	"fmt"
	"log"

	Configs "github.com/Lucasmartinsn/grocery-api/Configs/confDB"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func connectionType() (driver, sc string) {
	conf := Configs.GetDB()

	switch conf.Driver {
	case "postgres":
		driver = "postgres"
		sc = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.Host, conf.Port, conf.User, conf.Pass, conf.Database,
		)
	case "mysql":
		driver = "mysql"
		sc = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			conf.User, conf.Pass, conf.Host, conf.Port, conf.Database,
		)
	default:
		log.Fatalf("Driver desconhecido: %s", conf.Driver)
	}
	return
}

func OpenConnection() (*sql.DB, error) {
	driver, sc := connectionType()

	// Open Connection
	conn, err := sql.Open(driver, sc)
	if err != nil {
		log.Fatalf("error database connection: %v", err)
	}

	// Testing connection
	err = conn.Ping()
	return conn, err
}
