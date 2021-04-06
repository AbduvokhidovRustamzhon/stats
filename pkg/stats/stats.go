package stats

import (
	"github.com/AbduvokhidovRustamzhon/bank/v2/pkg/types"
)

//Avg - берет среднее значение
func Avg(payments []types.Payment) (money types.Money) {
	k := 0
	for _, payment := range payments {
	  if payment.Status != types.StatusFail {
		money += payment.Amount
		k++
	  }
	}
	return money / types.Money(k)
  }
  
  // TotalInCategory -
  func TotalInCategory(payments []types.Payment, category types.Category) (money types.Money) {
	for _, payment := range payments {
	  if payment.Category == category && payment.Status != types.StatusFail {
		money += payment.Amount
	  }
	}
	return
  }

  //TODO: считает средуню сумму
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	avg := map[types.Category]types.Money{}
	categories := map[types.Category]types.Money{}


	for _, value := range  payments{
		categories[value.Category] += value.Amount
		avg[value.Category] ++
	}
	// тут сперва сделали плюсы, теперь

	for key := range categories{
		categories[key] /= avg[key]
	}




	return categories

}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money, ) map[types.Category]types.Money {

	// 1 результат выводим
	result := map[types.Category]types.Money{}


	for key, value := range first{
		result[key] -= value
	}

	for key, value := range second{
		result[key] += value
	}



	return result


}