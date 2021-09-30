package stats

import "github.com/jkodirov/bank/v2/pkg/types"

func Avg(payments []types.Payment) types.Money {
	sum := 0
	count := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			sum += int(payment.Amount)
			count += 1
		}
	}
	avg := types.Money(sum / count)
	return avg
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Category == category {
			sum += int(payment.Amount)
		}
	}
	return types.Money(sum)
}

func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment

	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}