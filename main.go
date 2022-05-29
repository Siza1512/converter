package main

import "errors"

var negativeNumberNotSupport = errors.New("negative number not support")

func convertArshinToM(n float64) (float64, error) {
	if n < 0 {
		return 0, negativeNumberNotSupport
	}
	return n * 0.71, nil
}
