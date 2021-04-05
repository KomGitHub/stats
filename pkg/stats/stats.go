package stats

import (
	"github.com/KomGitHub/bank/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	amount := types.Money(0)
	notFailed := types.Money(0)
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			amount += payment.Amount
			notFailed += 1
		}
	}
	return amount / notFailed
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	amount := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		if payment.Status == types.StatusFail {
			continue
		}
		amount += payment.Amount
	}
	return amount
}

// FilterByCategory возвращает платежи в указанной категории.
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}

// CategoriesTotal возвращает сумму платежей по каждой категории.
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

// CategoriesAvg возвращает среднюю сумму платежей по каждой категории.
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categoriesSum := map[types.Category]types.Money{}
	categoriesCount := map[types.Category]types.Money{}
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categoriesSum[payment.Category] += payment.Amount
		categoriesCount[payment.Category] += 1
	}

	for key := range categoriesSum {
		categories[key] = categoriesSum[key] / categoriesCount[key]
	}

	return categories
}

// PeriodsDynamic сравнивает расходы по категориям за два периода.
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	
	for key := range first {
		categories[key] = second[key] - first[key]
	}

	for key := range second {
		if _, ok := first[key]; !ok {
			categories[key] = second[key]
		}
	}

	return categories
}