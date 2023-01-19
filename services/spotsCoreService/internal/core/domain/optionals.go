package domain

type Optional struct {
	value string
}

func NewOptional(value string) *Optional {
	return &Optional{
		value: value,
	}
}
func (op *Optional) IsPresent() bool {
	return op.value != ""
}

// TODO -> We should check that the value is not empty if not we should raise an error
func (op *Optional) GetValue() string {
	return op.value
}
