package stats

import "github.com/jkodirov/bank/pkg/types"

func Avg(payments []types.Payment) types.Money {
	sum := 0
	for _, payment := range payments {
		sum += int(payment.Amount)
	}
	avg := types.Money(sum / len(payments))
	return avg
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0
	for _, payment := range payments {
		if payment.Category == category {
			sum += int(payment.Amount)
		}
	}
	return types.Money(sum)
}