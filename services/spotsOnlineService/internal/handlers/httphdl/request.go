package httphdl

type SpotDateRequest struct {
	DurationApproximated int64 `json:"durationApproximated" binding:"required"`
	MaximunCapacity      int   `json:"maximunCapacity" binding:"required"`
}
