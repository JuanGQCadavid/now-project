package main

import (
	"errors"

	"github.com/JuanGQCadavid/now-project/services/scheduledPatternsChecker/internal/core/domain"
)

var (
	ErrEmptySpot = errors.New("err empty spot to transform")
)

func fromSpotRequestToSpot(spotRequest SpotRequest) domain.Spot {
	var spot domain.Spot

	spot.SpotId = spotRequest.SpotInfo.SpotId
	schedulePatterns := make([]domain.SchedulePattern, len(spotRequest.SpotPatterns))

	for index, sp := range spotRequest.SpotPatterns {
		schedulePatterns[index] = domain.SchedulePattern{
			Id:        sp.PatternId,
			HostId:    sp.SpotHost.HostId,
			Day:       domain.Day(sp.Day),
			FromDate:  sp.FromDate,
			ToDate:    sp.ToDate,
			StartTime: sp.StartTime,
			EndTime:   sp.EndTime,
		}
	}

	return spot
}
