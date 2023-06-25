package main

type Body struct {
	ScheduleId  string      `json:"scheduleId"`
	SpotId      string      `json:"spotId"`
	UserId      string      `json:"userId"`
	SpotRequest SpotRequest `json:"aditionalpayload"`
}

type SpotRequest struct {
	SpotInfo     SpotInfo       `json:"spotInfo"`
	SpotPatterns []SpotPatterns `json:"patterns"`
}

type SpotInfo struct {
	SpotId string `json:"spotId"`
}

type SpotPatterns struct {
	PatternId string   `json:"patternId"`
	SpotHost  SpotHost `json:"host"`
	SpotSate  SpotSate `json:"state"`
	Day       int      `json:"day"`
	FromDate  string   `json:"fromDate"`
	ToDate    string   `json:"toDate"`
	StartTime string   `json:"StartTime"`
	EndTime   string   `json:"endTime"`
}

type SpotSate struct {
	Status string `json:"status"`
	Since  int    `json:"since"`
}

type SpotHost struct {
	HostId string `json:"hostId"`
}
