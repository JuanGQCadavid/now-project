package commands

func getStringFromInterface(param interface{}) string {
	if param == nil {
		return ""
	}
	return param.(string)
}

func getFloat64FromInterface(param interface{}) float64 {
	if param == nil {
		return 0
	}
	return param.(float64)
}

func getInt64FromInterface(param interface{}) int64 {
	if param == nil {
		return 0
	}
	return param.(int64)
}
