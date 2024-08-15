package data

import (
	"database/sql"
	"fmt"
	"os"

	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

// var db *sql.DB
var upper db2.Session

type Models struct {
	// Any models inserted here (and in the new() function)
	// are easily accessible throughout the entire application
}

func New(databasePool *sql.DB) Models {
	//db = databasePool

	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mysqldb", "mariadb", "maria:":
		upper, _ = mysql.New(databasePool)
	case "postgres", "postgresql":
		upper, _ = postgresql.New(databasePool)
	default:
		// do nothing
	}

	return Models{}
}

func GetInsertID(i db2.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	}
	return int(i.(int32))
}
