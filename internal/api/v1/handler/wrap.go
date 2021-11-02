package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mabta/clpc/errs"
)

func JSONWrap(f func(*gin.Context) *errs.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := f(c); err != nil {
			log.Printf("[%s] %#v\n", c.Request.RequestURI, err)
			c.JSON(http.StatusOK, gin.H{
				"code": 111,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
	}
}
