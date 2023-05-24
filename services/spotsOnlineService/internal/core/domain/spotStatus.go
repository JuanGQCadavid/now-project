package domain

const (
	ONLINE_SPOT    SpotStatus = "ONLINE"
	PAUSED_SPOT    SpotStatus = "PAUSED"
	FINALIZED_SPOT SpotStatus = "FINALIZED"
	UNKNOWN        SpotStatus = "UNKNOWN"
)

type SpotStatus string
