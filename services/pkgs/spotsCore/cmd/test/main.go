package main

import (
	"github.com/JuanGQCadavid/now-project/services/spotsCore/internal/core/domain"
	spotactivityservices "github.com/JuanGQCadavid/now-project/services/spotsCore/internal/repositories/spotActivityServices"
)

func main() {
	location := spotactivityservices.NewAWSSpotActivityTopic()

	err := location.AppendSpot(domain.Spot{})
	panic(err)
}
