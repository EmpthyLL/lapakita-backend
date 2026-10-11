package utils

import "fmt"

// FormatRupiah formats a whole-number IDR amount with dot group separators.
func FormatRupiah(amount float64) string {
	value := fmt.Sprintf("%d", int64(amount))
	if len(value) <= 3 {
		return value
	}

	var formatted []byte
	for i, char := range value {
		if i > 0 && (len(value)-i)%3 == 0 {
			formatted = append(formatted, '.')
		}
		formatted = append(formatted, byte(char))
	}
	return string(formatted)
}
