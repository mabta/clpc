package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mabta/clpc/internal/db"
)

type PeriodResp struct {
	Period            int    `json:"period"`
	Numbers           string `json:"numbers"`
	OpenTime          string `json:"openTime"`
	CollectTime       string `json:"collectTime"`
	NextPeriod        int    `json:"nextPeriod"`
	NextPerIodTime    string `json:"nextPeriodTime"`
	DisNextPerIodTime int    `json:"disNextPeriodTime"`
}

func PeriodRespFromIssue(issue *db.Issue) *PeriodResp {
	return &PeriodResp{
		Period:      0,
		Numbers:     issue.Result,
		OpenTime:    issue.Schedule,
		CollectTime: "",
	}
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func FailedWithCode(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
func Failed(c *gin.Context, msg string) {
	FailedWithCode(c, 1201, msg)
}
