package rds

import (
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MysqlConnector struct {
	session *gorm.DB
}

func NewConector(dbUser string, dbPassword string, dbName string, dbUrl string) (*MysqlConnector, error) {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	url := fmt.Sprintf("postgres://%s:%s@%s:3306/%s?TimeZone=UTC", dbUser, dbPassword, dbUrl, dbName)
	session, err := gorm.Open(postgres.Open(url), &gorm.Config{})

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

func (conn *MysqlConnector) Migrate() {
	conn.session.AutoMigrate(&domain.DatesLocation{}, &domain.States{}, &domain.Types{})
}
