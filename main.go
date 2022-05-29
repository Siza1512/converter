package main

import "errors"

func convertArshinToM(n float64) (float64, error) {
	if n == -1 {
		return 0, errors.New("error")
	}
	return n * 0.71, nil
}
