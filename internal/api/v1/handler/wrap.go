package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mabta/clpc/errs"
	"github.com/mabta/clpc/internal/log"
)

func JSONWrap(f func(*gin.Context) *errs.Error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := f(c); err != nil {
			log.Logger.Errorf("[%s] %#v\n", c.Request.RequestURI, err)
			c.JSON(http.StatusOK, gin.H{
				"code": err.Code,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
	}
}
