package stats

import(
	"github.com/AbduvokhidovRustamzhon/bank/pkg/types"
)

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	avg := types.Money(0)
	for _, payment := range payments {
		avg += payment.Amount;
	}
	avg = avg/types.Money(len(payments))
	return avg
}

//TotalInCategory находит сумму покупок в заднной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	total := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category{
			total += payment.Amount
		}
	}
	return total
}