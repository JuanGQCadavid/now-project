package main

import (
	"github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/core/domain"
	spotactivityservices "github.com/JuanGQCadavid/now-project/services/spotsCoreService/internal/repositories/spotActivityServices"
)

func main() {
	location := spotactivityservices.NewAWSSpotActivityTopic()

	err := location.AppendSpot(domain.Spot{})
	panic(err)
}
