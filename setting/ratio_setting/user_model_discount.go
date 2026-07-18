package ratio_setting

import (
	"github.com/QuantumNous/new-api/types"
)

// 用户ID-{模型-折扣}
var userModelDiscounts = types.NewRWMap[int, map[string]int]()

func GetUserModelDiscount(userId int, modelName string) float64 {
	modelDiscounts, ok := userModelDiscounts.Get(userId)
	if !ok {
		return 1
	}
	discount, ok := modelDiscounts[modelName]
	if !ok {
		return 1
	}
	return float64(discount) / 10_000
}

func ReplaceUserModelDiscounts(userId int, modelDiscounts map[string]int) {
	if len(modelDiscounts) == 0 {
		userModelDiscounts.Set(userId, modelDiscounts)
		return
	}
	copiedModelDiscounts := make(map[string]int, len(modelDiscounts))
	for modelName, discount := range modelDiscounts {
		if modelName != "" && discount > 0 {
			copiedModelDiscounts[modelName] = discount
		}
	}
	if len(copiedModelDiscounts) == 0 {
		userModelDiscounts.Set(userId, copiedModelDiscounts)
		return
	}
	userModelDiscounts.Set(userId, copiedModelDiscounts)
}
