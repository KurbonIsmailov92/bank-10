package stats

import (
	"fmt"

	"github.com/KurbonIsmailov92/bank/pkg/bank/types"
)

func ExampleAvg() {

	payments := []types.Payment{
		{
			ID: 123,
			Amount: 10,
			Category: "food",
		},
		{
			ID: 124,
			Amount: 20,
			Category: "tax",
		},
		{
			ID: 125,
			Amount: 15,
			Category: "car",
		},
		{
			ID: 126,
			Amount: 10,
			Category: "food",
		},

	}

	result:= Avg(payments)

	fmt.Println(result)

	// Output: 13
}

func ExampleTotalInCategory() {

	payments := []types.Payment{
		{
			ID:       123,
			Amount:   10,
			Category: "food",
		},
		{
			ID:       124,
			Amount:   20,
			Category: "tax",
		},
		{
			ID:       125,
			Amount:   15,
			Category: "car",
		},
		{
			ID:       126,
			Amount:   10,
			Category: "food",
		},
	}

	result := TotalInCategory(payments, "food")

	fmt.Println(result)

	// Output: 20
}