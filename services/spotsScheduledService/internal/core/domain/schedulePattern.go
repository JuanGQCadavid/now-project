package domain

import (
	"errors"
	"time"

	"github.com/JuanGQCadavid/now-project/services/spotsScheduledService/internal/core/logs"
)

var (
	ErrParsingTimes = errors.New("We face an error while parsing the times")
)

const (
	Monday Day = 1 << iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var Days []Day = []Day{
	Monday,
	Tuesday,
	Wednesday,
	Thursday,
	Friday,
	Saturday,
	Sunday,
}

type Day uint

func IsMonday(day Day) bool {
	return (day & Monday) == Monday
}

func IsTuesday(day Day) bool {
	return (day & Tuesday) == Tuesday
}

func IsWednesday(day Day) bool {
	return (day & Wednesday) == Wednesday
}

func IsThursday(day Day) bool {
	return (day & Thursday) == Thursday
}

func IsFriday(day Day) bool {
	return (day & Friday) == Friday
}

func IsSaturday(day Day) bool {
	return (day & Saturday) == Saturday
}

func IsSunday(day Day) bool {
	return (day & Sunday) == Sunday
}

type SchedulePattern struct {
	Id        string `json:"patternId,omitempty"`
	State     State  `json:"state,omitempty"`
	Host      Host   `json:"host,omitempty"`
	Day       Day    `json:"day,omitempty"`
	FromDate  string `json:"fromDate,omitempty"`
	ToDate    string `json:"toDate,omitempty,omitempty"`
	StartTime string `json:"StartTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

// result["fromDate"] = fromDateTime
// result["toDate"] = toDateTime
// result["StartTime"] = StartTimeTime
// result["endTime"] = endTimeTime

func (sp *SchedulePattern) Overlaps(otherSchedulePattern SchedulePattern) (bool, error) {
	logs.Info.Printf("Overlaps, Self Id: %s, agains id: %s \n", sp.Id, otherSchedulePattern.Id)

	slefTimes, err := sp.fromStringToTime(sp.FromDate, sp.ToDate, sp.StartTime, sp.EndTime)

	if err != nil {
		return true, ErrParsingTimes
	}

	otherTimes, err := sp.fromStringToTime(otherSchedulePattern.FromDate, otherSchedulePattern.ToDate, otherSchedulePattern.StartTime, otherSchedulePattern.EndTime)

	if err != nil {
		return true, ErrParsingTimes
	}

	// If dates don't overlap then we don't need to check the rest
	if !sp.doesDatesOverlap(slefTimes["fromDate"], slefTimes["toDate"], otherTimes["fromDate"], otherTimes["toDate"]) {
		return false, nil
	}

	// As they overlap by dates we should go day by day and check it using the start and end time
	for _, day := range Days {
		if ((sp.Day & day) == day) && ((otherSchedulePattern.Day & day) == day) {
			// They share the same day, so we need to check the start and end time
			if sp.doesTimesOverlap(slefTimes["StartTime"], slefTimes["endTime"], otherTimes["StartTime"], otherTimes["endTime"]) {
				logs.Error.Println("Dates overlap and days match on a time fraction")
				return true, nil
			}
		}
	}

	logs.Info.Println("Dates DONT overlap")
	return false, nil
}

func (sp *SchedulePattern) doesTimesOverlap(StartA time.Time, EndA time.Time, StartB time.Time, EndB time.Time) bool {
	if StartA.Compare(StartB) < 1 {
		if StartB.Compare(EndA) < 1 {
			return true
		}
	} else {
		if StartA.Compare(EndB) < 1 {
			return true
		}
	}

	return false
}

func (sp *SchedulePattern) doesDatesOverlap(StartA time.Time, EndA time.Time, StartB time.Time, EndB time.Time) bool {
	//logs.Info.Printf("doesDatesOverlap: \n\tStartA: %+v, \n\tEndA: %+v, \n\tStartB: %+v, \n\tEndB: %+v \n", StartA, EndA, StartB, EndB)

	if StartA.Compare(StartB) < 1 {
		if StartB.Compare(EndA) < 1 {
			return true
		}
	} else {
		if StartA.Compare(EndB) < 1 {
			return true
		}
	}
	// if StartA.Before(StartB) {
	// 	if StartB.Before(EndA) {
	// 		return true
	// 	}
	// } else {
	// 	if StartA.Before(EndB) {
	// 		return true
	// 	}
	// }
	logs.Info.Println("Dates NOT overlap")
	return false
}

func (sp *SchedulePattern) fromStringToTime(fromDate string, toDate string, StartTime string, endTime string) (map[string]time.Time, error) {
	result := make(map[string]time.Time, 4)

	fromDateTime, err := time.Parse(time.DateOnly, fromDate)

	if err != nil {
		return nil, err
	}

	toDateTime, err := time.Parse(time.DateOnly, toDate)

	if err != nil {
		return nil, err
	}
	StartTimeTime, err := time.Parse(time.TimeOnly, StartTime)

	if err != nil {
		return nil, err
	}
	endTimeTime, err := time.Parse(time.TimeOnly, endTime)

	if err != nil {
		return nil, err
	}

	result["fromDate"] = fromDateTime
	result["toDate"] = toDateTime
	result["StartTime"] = StartTimeTime
	result["endTime"] = endTimeTime

	//logs.Info.Printf("fromStringToTime: \n\tfromDate: %+v, \n\ttoDate: %+v, \n\tStartTime: %+v, \n\tendTime: %+v \n", result["fromDate"], result["toDate"], result["StartTime"], result["endTime"])
	return result, nil
}
