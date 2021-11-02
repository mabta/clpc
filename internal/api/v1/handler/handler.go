package handler

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mabta/clpc/errs"
	"github.com/mabta/clpc/internal/db"
)

// Current 当前期数的结果
func Current(c *gin.Context) *errs.Error {
	lotteryType := c.Query("lotteryType")
	if lotteryType == "" {
		return errs.NewInvalidParam()
	}
	issues, err := db.TopIssues(1, "ticket=$1", lotteryType)
	if err != nil {
		return errs.ToDbError(err)
	}
	if issues == nil || len(issues) == 0 {
		return errs.NotFoundError()
	}
	Success(c, PeriodRespFromIssue(issues[0]))
	return nil
}

// Latest 最新结果
func Latest(c *gin.Context) *errs.Error {
	lotteryType := c.Query("lotteryType")
	if lotteryType == "" {
		return errs.NewInvalidParam()
	}
	issues, err := db.TopIssues(30, "ticket=$1", lotteryType)
	if err != nil {
		return errs.ToDbError(err)
	}
	resp := make([]*PeriodResp, 0, len(issues))
	for _, issue := range issues {
		resp = append(resp, PeriodRespFromIssue(issue))
	}
	Success(c, resp)
	return nil
}

// Period 某一期
func Period(c *gin.Context) *errs.Error {
	lotteryType := c.Query("lotteryType")
	period := c.Query("period")
	if lotteryType == "" || period == "" {
		return errs.NewInvalidParam()
	}
	issue, err := db.FindIssueBy("ticket=$1 AND period=$2", lotteryType, period)
	if err == sql.ErrNoRows {
		return errs.NotFoundError()
	}
	if err != nil {
		return errs.ToDbError(err)
	}
	Success(c, PeriodRespFromIssue(issue))
	return nil
}

// 服务器时间
func Now(c *gin.Context) *errs.Error {
	Success(c, gin.H{"now": time.Now().Format("2006-01-02 15:04:05")})
	return nil
}
