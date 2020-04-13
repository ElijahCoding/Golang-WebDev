package handler

import (
	"newsfeeder/platform/newsfeed"
	"net/http"
	"github.com/gin-gonic/gin"
)

func NewsfeedPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
            "hello": "hi",
        })
	}
}
