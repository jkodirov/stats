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

// returns sum of payments per category
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	for _, payment :=range payments {
		categories[payment.Category] += payment.Amount
	}
	return categories
}

// avg payment amount per category
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	categoriesCount := map[types.Category]int{}
	for _, payment :=range payments {
		categories[payment.Category] += payment.Amount
		categoriesCount[payment.Category] += 1
	}
	for key, value := range categoriesCount {
		categories[key] /= types.Money(value)
	}
	return categories
}


func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	for key := range first {
		second[key] -= first[key]
	}
	return second
}