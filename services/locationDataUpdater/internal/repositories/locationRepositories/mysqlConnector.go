package locationrepositories

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnector struct {
	conectorType string
	dbUser       string
	dbPassword   string
	dbName       string
	dbUrl        string
	session      *sql.DB
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

func NewConectorFromEnv() *MysqlConnector {
	dbUser, isPresentUser := os.LookupEnv("dbUser")
	dbPassword, isPresentPass := os.LookupEnv("dbPassword")
	dbName, isPresentName := os.LookupEnv("dbName")
	dbUrl, isPresentUrl := os.LookupEnv("dbUrl")

	if !isPresentUrl || !isPresentName || !isPresentPass || !isPresentUser {
		log.Println("dbUser: ", dbUser)
		log.Println("dbPassword: ", dbPassword)
		log.Println("dbName: ", dbName)
		log.Println("dbUrl: ", dbUrl)
		log.Fatalln("The ULR, Password or Username, dbName is not present in the env.")
	}

	return NewConector("mysql", dbUser, dbPassword, dbName, dbUrl)
}

func (con *MysqlConnector) createSession() (*sql.DB, error) {
	return sql.Open(con.conectorType, fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", con.dbUser, con.dbPassword, con.dbUrl, con.dbName))
}

func (con *MysqlConnector) GetSession() (*sql.DB, error) {

	if con.session == nil {
		logs.Info.Println("Creating session")
		session, err := con.createSession()

		if err != nil {
			logs.Error.Println("We fail creating the session, error: ", err.Error())
			return nil, err
		}

		con.session = session
	}

	return con.session, nil
}
