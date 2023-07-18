package domain

type Date struct {
	StartTime string `json:"startTime,omitempty"`
	Confirmed bool   `json:"confirmed"`
	Id        string `json:"id,omitempty"`
	SpotId    string `json:"spotId,omitempty"`
	OnPlace   Place  `json:"place,omitempty"`
	DateStamp string `json:"dateStamp,omitempty"`
	Host      Host   `json:"host,omitempty"`
}

type Host struct {
	HostId   string `json:"hostId,omitempty"`
	HostName string `json:"hostName,omitempty"`
}

type Place struct {
	Name          string  `json:"name,omitempty"`
	Lat           float64 `json:"lat,omitempty"`
	Lon           float64 `json:"lon,omitempty"`
	MapProviderId string  `json:"mapProviderId,omitempty"`
}
