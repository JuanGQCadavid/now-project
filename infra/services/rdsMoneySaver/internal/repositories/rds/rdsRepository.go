package rds

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

const (
	DB_ALREADY_UP   string = "InvalidDBInstanceState: Instance locations-db is not in stopped"
	DB_ALREADY_DOWN string = "InvalidDBInstanceState: Instance locations-db is not in available state"
)

type RDSRepository struct {
	service *rds.RDS
}

func NewRDSRepository() *RDSRepository {

	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	service := rds.New(session)
	return &RDSRepository{
		service: service,
	}
}

func (repo *RDSRepository) StopInstance(instanceId string) error {
	log.Println("StopInstance - ", instanceId)
	input := &rds.StopDBInstanceInput{
		DBInstanceIdentifier: aws.String(instanceId),
	}
	out, err := repo.service.StopDBInstance(input)

	if err != nil {
		if strings.Contains(err.Error(), DB_ALREADY_DOWN) {
			log.Println("INFO: DB already down.")
			return nil
		} else {
			log.Println("ERROR: ", err.Error())
			return err
		}
	}

	log.Println(fmt.Sprintf("OUTPUT: %+v", out))

	return nil
}

func (repo *RDSRepository) StartInstance(instanceId string) error {
	log.Println("StartInstance - ", instanceId)

	input := &rds.StartDBInstanceInput{
		DBInstanceIdentifier: aws.String(instanceId),
	}
	out, err := repo.service.StartDBInstance(input)

	if err != nil {

		if strings.Contains(err.Error(), DB_ALREADY_UP) {
			log.Println("INFO: DB already on.")
			return nil
		} else {
			log.Println("ERROR: ", err.Error())
			return err
		}
	}

	log.Println(fmt.Sprintf("OUTPUT: %+v", out))

	return nil
}
