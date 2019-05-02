package helpers

import "strconv"

// StringToInt string to int
func StringToInt(char string) (int, error) {
	convert, err := strconv.Atoi(char)
	if err != nil {
		return 0, err
	}
	return convert, nil
}
