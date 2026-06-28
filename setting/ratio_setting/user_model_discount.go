package ratio_setting

import "sync"

var (
	userModelDiscountMu sync.RWMutex
	userModelDiscounts  = map[int]map[string]int{}
)

func GetUserModelDiscount(userId int, modelName string) float64 {
	userModelDiscountMu.RLock()
	defer userModelDiscountMu.RUnlock()

	modelDiscounts, ok := userModelDiscounts[userId]
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
	userModelDiscountMu.Lock()
	defer userModelDiscountMu.Unlock()

	if len(modelDiscounts) == 0 {
		delete(userModelDiscounts, userId)
		return
	}

	copiedModelDiscounts := make(map[string]int, len(modelDiscounts))
	for modelName, discount := range modelDiscounts {
		if modelName != "" && discount > 0 {
			copiedModelDiscounts[modelName] = discount
		}
	}
	if len(copiedModelDiscounts) == 0 {
		delete(userModelDiscounts, userId)
		return
	}
	userModelDiscounts[userId] = copiedModelDiscounts
}
