package main

import (
	"fmt"
	"universal-converter/converters"
)

func main() {
	var n float64
	fmt.Println("Введите аршины")
	_, _ = fmt.Scanf("%f", &n)
	m, err := converters.ConvertArshinToM(n)
	if err != nil {
		fmt.Println("не удалось перевести аршины в метры err:", err)
		return
	}
	fmt.Printf("%.2f аршин равно %.2f метров", n, m)
}
