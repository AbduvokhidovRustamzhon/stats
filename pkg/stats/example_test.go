package stats

import (
	"fmt"

	"github.com/AbduvokhidovRustamzhon/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 5000_00,
		},
		{
			Amount: 0,
		},
	}
	fmt.Println(Avg(payments))
	//Output:
	//250000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount: 5000_00,
			Category: "tech",
		},
		{
			Amount: 0,
			Category: "no tech",
		},
	}
	fmt.Println(TotalInCategory(payments,types.Category("tech")))
	//Output:
	//500000
}
