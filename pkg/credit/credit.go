package credit

/*
sumCredit = сумма кредита
periodCredit = срок, на который берется кредит
rateCredit = ставка процентная кредита
*/

import (
	"math"
)

// Calculate the credit
func Calculate(sumCredit, periodCredit, rateCredit int64) (monthPayment, sumPayment, overPayment int64) {
	// Ежемесячная процентная ставка
	i := float64(rateCredit) / 12 / 100
	// fmt.Printf("Ежемесячная процентная ставка: %0.4v\n", i)

	// Коэфицент аннуитета
	k := i * math.Pow((1+i), float64(periodCredit)) / (math.Pow((1+i), float64(periodCredit)) - 1)
	// fmt.Printf("Коэфицент аннуитета: %.4v\n", k)

	a := k * float64(sumCredit)
	monthPayment = int64(a) * 100
	sumPayment = (periodCredit * monthPayment) * 100
	overPayment = (sumPayment - sumCredit) * 100

	return
}
