package db

import (
	"time"

	"github.com/mabta/clpc/eth/blocktime"
)

// Issue 开奖结果
type Issue struct {
	ID          uint64
	Ticket      string
	Result      string
	Schedule    string
	BlockTime   uint64
	BlockNumber uint64
	Dateline    uint64
	DateStr     string
}

func NewIssue(ticket, result, schedule string, blockTime, blockNumber uint64) *Issue {
	return &Issue{
		Ticket:      ticket,
		Result:      result,
		Schedule:    schedule,
		BlockTime:   blockTime,
		BlockNumber: blockNumber,
		Dateline:    uint64(time.Now().Unix()),
		DateStr:     blocktime.DateStr(blockTime),
	}
}

// InsertIssue 插入开奖结果
func InsertIssue(issue *Issue) (uint64, error) {
	stmt, err := db.Prepare("INSERT INTO issue (ticket, result, schedule, block_time, block_number, dateline, date_str) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(issue.Ticket, issue.Result, issue.Schedule, issue.BlockTime, issue.BlockNumber, issue.Dateline, issue.DateStr)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}
