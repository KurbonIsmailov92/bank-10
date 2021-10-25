package stats

import "github.com/KurbonIsmailov92/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {

	sum := types.Money(0)
	count := int(0)

	for _, payment := range payments {

		sum += payment.Amount
	}

	count = len(payments)

	return sum / types.Money(count)

}