package main

import (
	"fmt"

	"github.com/sorokinkir/bgo2.1/pkg/credit"
)

func main() {
	monthPayment, sumPayment, overPayment := credit.Calculate(1_000_000, 36, 20)
	
	fmt.Printf("Ежемесячный платеж: %v\n", monthPayment)
	fmt.Printf("Переплата по кредиту: %v\n", overPayment)
	fmt.Printf("Общая выплата: %v\n", sumPayment)
}
