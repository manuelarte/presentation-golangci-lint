package middlewares

import (
	"github.com/gin-gonic/gin"
)

func NewPaginationWithDefaultSize(defaultSize int) func(c *gin.Context) {
	return pagination(defaultSize)
}

func pagination(defaultSize int) func(c *gin.Context) {
	return func(c *gin.Context) {
		size, ok := c.GetQuery("size")
		if ok {
			c.Set("size", size)
		} else {
			c.Set("size", defaultSize)
		}

		c.Next()
	}
}
