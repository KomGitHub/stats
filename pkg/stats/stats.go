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