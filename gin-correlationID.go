package ginCorrelationID

import (
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/nu7hatch/gouuid"
)

//CorrelationIDMiddleware adds correlationID if it's not specified in HTTP request
func CorrelationIDMiddleware() gin.HandlerFunc {
	return addCorrelationID
}

func addCorrelationID(c *gin.Context) {
	corralationID := c.Request.Header.Get("CorrelationID")

	if strings.TrimSpace(corralationID) == "" {
		id, _ := uuid.NewV4()

		c.Request.Header.Add("CorrelationID", id.String())
	}

	c.Next()
}
