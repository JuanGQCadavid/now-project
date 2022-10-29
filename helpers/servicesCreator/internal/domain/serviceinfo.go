package domain

import "strings"

type ServiceName string

func (s *ServiceName) ToCapitalCase() string {
	oldValue := string(*s)
	if len(oldValue) == 0 {
		return oldValue
	}

	firstLetter := oldValue[0]

	if firstLetter >= 65 && firstLetter <= 90 {
		return oldValue
	} else if firstLetter >= 97 && firstLetter <= 122 {
		firstLetter = firstLetter - 32
		return strings.Replace(oldValue, string(oldValue[0]), string(firstLetter), 1)
	} else {
		return oldValue
	}
}

func (s *ServiceName) ToLoweCase() string {
	oldValue := string(*s)
	if len(oldValue) == 0 {
		return oldValue
	}

	firstLetter := oldValue[0]

	if firstLetter >= 65 && firstLetter <= 90 {
		firstLetter = firstLetter + 32
		return strings.Replace(oldValue, string(oldValue[0]), string(firstLetter), 1)
	} else if firstLetter >= 97 && firstLetter <= 122 {
		return oldValue

	} else {
		return oldValue
	}
}

type ServiceInformation struct {
	ServicePath string
	ServiceName ServiceName
	// ServiceLambdaNale string
}

func NewServiceInformation(servicePath string, serviceName ServiceName) *ServiceInformation {
	return &ServiceInformation{
		ServicePath: servicePath,
		ServiceName: serviceName,
		// ServiceLambdaNale: serviceLambdaNale,
	}
}
