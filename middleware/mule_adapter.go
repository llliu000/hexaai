package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MuleRequestConvert() func(c *gin.Context) {
	return func(c *gin.Context) {
		model := c.Param("model")
		if model == "" {
			abortWithOpenAiMessage(c, http.StatusBadRequest, "model parameter is required")
			return
		}

		c.Request.URL.Path = "/v1/video/generations"

		// Handle Jimeng official API request
		var originalReq map[string]any
		if err := json.NewDecoder(c.Request.Body).Decode(&originalReq); err != nil {
			c.Next()
			return
		}

		unifiedReq := map[string]interface{}{
			"model":    model,
			"prompt":   "mule",
			"metadata": originalReq,
		}

		jsonData, err := json.Marshal(unifiedReq)
		if err != nil {
			c.Next()
			return
		}

		// Update request body
		c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonData))

		c.Next()
	}
}
