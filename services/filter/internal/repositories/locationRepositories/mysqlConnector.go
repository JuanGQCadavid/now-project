package locationrepositories

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnector struct {
	conectorType string
	dbUser       string
	dbPassword   string
	dbName       string
	dbUrl        string
}

func NewConector(conectorType string, dbUser string, dbPassword string, dbName string, dbUrl string) *MysqlConnector {
	return &MysqlConnector{
		conectorType: conectorType,
		dbUser:       dbUser,
		dbPassword:   dbPassword,
		dbName:       dbName,
		dbUrl:        dbUrl,
	}

}

func (con *MysqlConnector) CreateSession() (*sql.DB, error) {
	return sql.Open(con.conectorType, fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", con.dbUser, con.dbPassword, con.dbUrl, con.dbName))
}
