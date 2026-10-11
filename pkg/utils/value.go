package utils

// GetIntVal returns zero for a nil integer pointer.
func GetIntVal(value *int) int {
	if value == nil {
		return 0
	}
	return *value
}

// GetFloatVal returns zero for a nil float pointer.
func GetFloatVal(value *float64) float64 {
	if value == nil {
		return 0
	}
	return *value
}
