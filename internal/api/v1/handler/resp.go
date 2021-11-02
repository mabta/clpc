package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mabta/clpc/errs"
	"github.com/mabta/clpc/eth/blocktime"
	"github.com/mabta/clpc/internal/db"
)

// PeriodResp 开奖结果响应
type PeriodResp struct {
	// Period 期数
	Period uint64 `json:"period"`
	// Numbers 开奖结果
	Numbers string `json:"numbers"`
	// OpenTime 计划开奖时间
	OpenTime string `json:"openTime"`
	// CollectTime 实际开奖时间
	CollectTime string `json:"collectTime"`
	// NextPeriod 下一期的期数
	NextPeriod uint64 `json:"nextPeriod"`
	// NextPerIodTime 下一期计划开奖时间
	NextPerIodTime string `json:"nextPeriodTime"`
	// DisNextPerIodTime 距离下一期开奖还有多久
	DisNextPerIodTime uint64 `json:"disNextPeriodTime"`
	// HasNextPeriod 是否有下期
	//HasNextPeriod bool `json:"hasNextPeriod"`
}

// PeriodRespFromIssue 将数据库中的开奖结果转换为开奖结果响应
func PeriodRespFromIssue(issue *db.Issue) *PeriodResp {
	return &PeriodResp{
		Period:            issue.Period,
		Numbers:           issue.Result,
		OpenTime:          blocktime.DateTimeStrToApi(issue.Schedule),
		CollectTime:       blocktime.DateTimeStrToApi(issue.BlockTime),
		NextPeriod:        issue.NextPeriod,
		NextPerIodTime:    blocktime.DateTimeStrToApi(issue.NextPeriodSchedule),
		DisNextPerIodTime: issue.NextPeriodSchedule - uint64(time.Now().Unix()),
	}
}

// Success 返回成功时的响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

// FailedWithCode 返回失败时的响应，并指定错误码
func FailedWithCode(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

// Failed 返回失败时的响应，使用默认错误码
func Failed(c *gin.Context, msg string) {
	FailedWithCode(c, errs.BadRequest, msg)
}
