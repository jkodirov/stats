package stats

import (
	"fmt"
	"github.com/jkodirov/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 120,
		},
		{
			ID: 2,
			Amount: 100,
		},
		{
			ID: 3,
			Amount: 80,
		},
	}
	fmt.Println(Avg(payments))
	//Output:
	// 100
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 120,
			Category: "Income",
		},
		{
			ID: 2,
			Amount: 100,
			Category: "Income",
		},
		{
			ID: 3,
			Amount: 80,
			Category: "Expense",
		},
	}
	fmt.Println(TotalInCategory(payments, "Income"))
	fmt.Println(TotalInCategory(payments, "Expense"))
	//Output:
	// 220
	// 80
}