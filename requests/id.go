package requests

import (
	"github.com/gin-gonic/gin"
)

const (
	RequestId         = "Request-Id"
	FallbackRequestID = "BEEFBAD-df50-11ed-9986-bff1b842c2a"
)

func ID(c *gin.Context) string {
	reqID, ok := c.Get(RequestId)
	if !ok {
		reqID = FallbackRequestID
	}
	return reqID.(string)
}
