package rds

import (
	"fmt"

	"github.com/JuanGQCadavid/now-project/services/locationDataUpdater/internal/core/domain"
	"github.com/JuanGQCadavid/now-project/services/pkgs/common/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_PORT int = 5432
)

type MysqlConnector struct {
	session *gorm.DB
}

func NewConector(dbUser string, dbPassword string, dbName string, dbUrl string) (*MysqlConnector, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?TimeZone=UTC", dbUser, dbPassword, dbUrl, DB_PORT, dbName)
	session, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		logs.Error.Println("We fail to create the connection to the DB, error: ", err.Error())
	}
	return &MysqlConnector{
		session: session,
	}, nil

}

func (conn *MysqlConnector) Migrate() {
	logs.Info.Println("Migrate command")
	conn.session.AutoMigrate(&domain.DatesLocation{}, &domain.States{}, &domain.Types{})
}
