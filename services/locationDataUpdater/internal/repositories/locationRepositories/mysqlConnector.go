package locationrepositories

import (
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConnector struct {
	session *gorm.DB
}

func NewConector(dbUser string, dbPassword string, dbName string, dbUrl string) (*MysqlConnector, error) {
	session, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", dbUser, dbPassword, dbUrl, dbName)), &gorm.Config{})

	if err != nil {
		logs.Error.Println("We fail to create the connection to the DB, error: ", err.Error())
	}
	return &MysqlConnector{
		session: session,
	}, nil

}

func NewConectorFromEnv() (*MysqlConnector, error) {
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

	return NewConector(dbUser, dbPassword, dbName, dbUrl)
}
