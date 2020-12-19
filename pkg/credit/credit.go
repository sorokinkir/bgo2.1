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
	percent := float64(rateCredit) / 12 / 100

	// Коэфицент аннуитета
	rate := percent * math.Pow((1+percent), float64(periodCredit)) / (math.Pow((1+percent), float64(periodCredit)) - 1)

	result := rate * float64(sumCredit)
	monthPayment = int64(result)
	sumPayment = (periodCredit * monthPayment)
	overPayment = (sumPayment - sumCredit)

	return convertToPenny(monthPayment), convertToPenny(overPayment), convertToPenny(sumPayment)
}

// convert penny from rubles
func convertToPenny(sum int64) int64 {
	return sum * 100
}