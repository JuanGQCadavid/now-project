package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/JuanGQCadavid/now-project/infra/services/rdsMoneySaver/internal/repositories/rds"
)

type MoneySaver struct {
	rdsRepo rds.RDSRepository
}

func NewMoneySaver(rdsRepo rds.RDSRepository) *MoneySaver {
	return &MoneySaver{
		rdsRepo: rdsRepo,
	}
}

func (ms *MoneySaver) StopDatabases(databases []string) error {
	log.Println(fmt.Sprintf("StopDatabases: %+v", databases))
	errCounter := 0
	for _, identifier := range databases {
		if len(identifier) > 0 {
			err := ms.rdsRepo.StopInstance(identifier)
			if err != nil {
				errCounter++
				log.Println("Error: ", err.Error())
			}
		}
	}

	if errCounter > 0 {
		return errors.New("There were an error while Stoping the instances, please check logs")
	}

	return nil
}

func (ms *MoneySaver) StartDatabases(databases []string) error {
	log.Println(fmt.Sprintf("StartDatabases: %+v", databases))
	errCounter := 0
	for _, identifier := range databases {
		if len(identifier) > 0 {
			err := ms.rdsRepo.StartInstance(identifier)
			if err != nil {
				errCounter++
				log.Println("Error: ", err.Error())
			}
		}
	}

	if errCounter > 0 {
		return errors.New("There were an error while Starting the instances, please check logs")
	}

	return nil
}
