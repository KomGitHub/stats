package stats

import (
	"github.com/KomGitHub/bank/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	amount := types.Money(0)
	for _, payment := range payments {
		amount += payment.Amount
	}
	return amount / types.Money(len(payments))
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	amount := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			amount += payment.Amount
		}
	}
	return amount
}