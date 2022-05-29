package converters

import "errors"

var NegativeNumberNotSupport = errors.New("negative number not support")

func ConvertArshinToM(n float64) (float64, error) {
	if n < 0 {
		return 0, NegativeNumberNotSupport
	}
	return n * 0.71, nil
}
