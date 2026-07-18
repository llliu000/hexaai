package codex

import (
	"slices"

	"github.com/QuantumNous/new-api/setting/ratio_setting"
)

var baseModelList = []string{}

var ModelList = slices.DeleteFunc(
	ratio_setting.WithCompactModelVariants(baseModelList),
	func(modelName string) bool {
		return modelName == ratio_setting.WithCompactModelSuffix("codex-auto-review")
	},
)

const ChannelName = "codex"
