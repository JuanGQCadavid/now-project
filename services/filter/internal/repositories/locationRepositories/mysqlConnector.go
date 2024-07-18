package locationrepositories

import (
	"fmt"
	"log"
	"os"

	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConector(dbUser string, dbPassword string, dbName string, dbUrl string) (*gorm.DB, error) {
	// session, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true", dbUser, dbPassword, dbUrl, dbName)), &gorm.Config{})
	url := fmt.Sprintf("postgres://%s:%s@%s:3306/%s?TimeZone=UTC", dbUser, dbPassword, dbUrl, dbName)
	session, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		logs.Error.Println("We fail to create the connection to the DB, error: ", err.Error())
	}
	return session, nil

}

func NewConectorFromEnv() (*gorm.DB, error) {
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

func Migrate(session *gorm.DB) {
	session.AutoMigrate(&DatesLocation{}, &States{}, &Types{})
}
