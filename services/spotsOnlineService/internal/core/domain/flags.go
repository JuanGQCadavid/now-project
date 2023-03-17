package domain

type SpotStateFlags uint

const (
	FlagOnline = 1 << iota
	FlagPaused
	FlagFinalized
)
