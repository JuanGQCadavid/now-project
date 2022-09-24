package main

import "github.com/JuanGQCadavid/now-project/infra/services/rdsMoneySaver/internal/repositories/rds"

func main() {
	dbInstance := "locations-db"
	rdsRepo := rds.NewRDSRepository()

	rdsRepo.StopInstance(dbInstance)
}
