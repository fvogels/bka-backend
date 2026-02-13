package util

import "strconv"

func ParseInt(str string) (int, error) {
	result, err := strconv.ParseInt(str, 10, 32)

	if err != nil {
		return 0, err
	}

	return int(result), nil
}
