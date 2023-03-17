package httphdl

type SpotDateRequest struct {
	DurationApproximated int64 `json:"durationApproximated" binding:"required"`
	MaximunCapacity      int64 `json:"maximunCapacity" binding:"required"`
}
