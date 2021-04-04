package stats

import (
	"github.com/KomGitHub/bank/pkg/types"
	"fmt"
)

func ExampleAvg() {
	fmt.Println(Avg([]types.Payment{
		{
			ID: 1,
			Amount: 1_000_00,
			Category: "Auto",
		},
	}))
	fmt.Println(Avg([]types.Payment{
		{
			ID: 1,
			Amount: 1_000_00,
			Category: "Auto",
		},
		{
			ID: 2,
			Amount: 3_000_00,
			Category: "Pharmacy",
		},
	}))
	//Output:
	//100000
	//200000
}

func ExampleTotalInCategory() {
	fmt.Println(TotalInCategory([]types.Payment{
		{
			ID: 1,
			Amount: 1_000_00,
			Category: "Auto",
		},
	}, "Pharmacy"))
	fmt.Println(TotalInCategory([]types.Payment{
		{
			ID: 1,
			Amount: 1_000_00,
			Category: "Auto",
		},
		{
			ID: 2,
			Amount: 3_000_00,
			Category: "Pharmacy",
		},
	}, "Pharmacy"))
	//Output:
	//0
	//300000
}