package stats

import (
	"github.com/AbduvokhidovRustamzhon/bank/v2/pkg/types"
	"reflect"
	"testing"
)

func TestCategoriesAvg_Multiple(t *testing.T) {
	payments := []types.Payment{
		{
			Amount: 1500,
			Category: "auto",
		},
		{
			Amount: 1500,
			Category: "mobi",
		},
		{
			Amount: 1500,
			Category: "mobi",
		},
		{
			Amount: 1500,
			Category: "auto",
		},


	}
	avg := CategoriesAvg(payments)
	expected := map[types.Category]types.Money{
		"auto": 1500,
		"mobi": 1500,
		}
		if !reflect.DeepEqual(expected, avg){
			t.Errorf("invalid result, expected: %v, actual: %v", expected, avg )
		}

}

func TestCategoriesAvg_empty(t *testing.T) {
	payments := []types.Payment{}
	avg := CategoriesAvg(payments)
	expected := map[types.Category]types.Money{}
	if !reflect.DeepEqual(expected, avg){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, avg )
	}

}


func TestCategoriesAvg_foundOne(t *testing.T) {
	payments := []types.Payment{
		{
			Amount: 1500,
			Category: "auto",
		},
		{
			Amount: 2000,
			Category: "auto",
		},
		{
			Amount: 2500,
			Category: "auto",
		},

	}
	avg := CategoriesAvg(payments)
	expected := map[types.Category]types.Money{
		"auto": 2000,
	}

	if !reflect.DeepEqual(expected, avg){
		t.Errorf("invalid result in test foundOne, expected: %v, actual: %v", expected, avg )
	}
}

func TestPeriodsDynamic_Multiple(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 1000,
		"mobi": 1000,
	}
	second := map[types.Category]types.Money{
		"auto": 500,
		"mobi": 500,
	}
	result := PeriodsDynamic(first, second)
	expected := map[types.Category]types.Money{
		"auto": -500,
		"mobi": -500,
}
	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result in test PeriodsDynamic_Multiple	, expected: %v, actual: %v", expected, result )
	}


}

func TestPeriodsDynamic_empty(t *testing.T) {
	first := map[types.Category]types.Money{}
	second := map[types.Category]types.Money{}
	result := PeriodsDynamic(first, second)
	expected := map[types.Category]types.Money{}
	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result in PeriodsDynamic_empty, expected: %v, got: %v", expected, result)
	}
}


func TestPeriodsDynamic_foundOne(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 1000,
	}
	second := map[types.Category]types.Money{
		"auto": 500,
	}
	result := PeriodsDynamic(first, second)
	expected := map[types.Category]types.Money{
		"auto": -500,
	}
	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result in PeriodsDynamic_foundOne, expected: %v, got: %v", expected, result)
	}
}