package middleware

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

func DoubaoRequestConvert() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Request.URL.Path = "/v1/video/generations" // 预先处理,防止在Distribute方法中出错

		var originalReq map[string]interface{}
		// common.UnmarshalBodyReusable 方法会做缓存处理，参数转换阶段暂时不做，由后面方法处理
		if err := json.NewDecoder(c.Request.Body).Decode(&originalReq); nil != err {
			c.Next()
			return
		}
		// Support both model_name and model fields
		model, _ := originalReq["model"].(string)

		unifiedReq := map[string]interface{}{
			"model":    model,
			"prompt":   "doubao", // 处理验证问题，后面直接使用metadata处理
			"metadata": originalReq,
		}

		jsonData, err := json.Marshal(unifiedReq)
		if err != nil {
			c.Next()
			return
		}

		// Rewrite request body and path
		c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonData))
		c.Next()
	}
}
