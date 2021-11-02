package handler

import "github.com/gin-gonic/gin"

func RegisterRouter(e *gin.Engine) {
	g := e.Group("/")
	{
		g.GET("/now/first", JSONWrap(Current))
		g.GET("/list", JSONWrap(Latest))
		g.GET("/period", JSONWrap(Period))
		g.GET("/now/time", JSONWrap(Now))
	}
}
