package db

import (
	"strconv"
	"strings"
	"time"

	"github.com/mabta/clpc/eth/blocktime"
)

// Issue 开奖结果
type Issue struct {
	ID uint64
	// Ticket 玩法
	Ticket string
	// Result 结果
	Result string
	// Schedule 计划开奖时间
	Schedule string
	// BlockTime 区块时间
	BlockTime uint64
	// BlockNumber 区块高度
	BlockNumber uint64
	// Dateline 入库时间
	Dateline uint64
	// DateStr 开奖日期字符表示
	DateStr string
	// BlockHash 区块哈希
	BlockHash string
}

func NewIssue(ticket, result, schedule, blockHash string, blockTime, blockNumber uint64) *Issue {
	return &Issue{
		Ticket:      ticket,
		Result:      result,
		Schedule:    schedule,
		BlockTime:   blockTime,
		BlockNumber: blockNumber,
		Dateline:    uint64(time.Now().Unix()),
		DateStr:     blocktime.DateStr(blockTime),
		BlockHash:   blockHash,
	}
}

// InsertIssue 插入开奖结果
func InsertIssue(issue *Issue) (uint64, error) {
	stmt, err := db.Prepare("INSERT INTO issue (ticket, result, schedule, block_time, block_number, dateline, date_str, block_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(issue.Ticket, issue.Result, issue.Schedule, issue.BlockTime, issue.BlockNumber, issue.Dateline, issue.DateStr, issue.BlockHash)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// FindIssueBy 根据指定条件查找单条开奖结果
func FindIssueBy(condition string, args ...interface{}) (*Issue, error) {
	sb := strings.Builder{}
	sb.WriteString("SELECT id,ticket, result, schedule, block_time, block_number, dateline, date_str,block_hash  FROM issue ")
	if condition != "" {
		sb.WriteString(" WHERE ")
		sb.WriteString(condition)
	}
	sb.WriteString(" LIMIT 1")
	stmt, err := db.Prepare(sb.String())
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	issue := new(Issue)
	row := stmt.QueryRow(args...)
	if err := row.Scan(&issue.ID, &issue.Ticket, &issue.Result, &issue.Schedule, &issue.BlockTime, &issue.BlockNumber, &issue.Dateline, &issue.DateStr, &issue.BlockHash); err != nil {
		return nil, err
	}
	return issue, nil
}

// FindIssue 根据ID查找单条开奖结果
func FindIssue(id uint64) (*Issue, error) {
	return FindIssueBy("id=$1", id)
}

// TopIssues 最新N期
func TopIssues(n int, condition string, args ...interface{}) ([]*Issue, error) {
	issues, _, err := SelectIssues(0, n, condition, args...)
	if err != nil {
		return nil, err
	}
	return issues, nil
}

// SelectIssues 分页查找所有符合条件的开奖结果
func SelectIssues(page, pageSize int, condition string, args ...interface{}) ([]*Issue, *Pagination, error) {
	sb := strings.Builder{}
	sbCount := strings.Builder{}

	sb.WriteString("SELECT id,ticket, result, schedule, block_time, block_number, dateline, date_str,block_hash  FROM issue ")
	sbCount.WriteString("SELECT COUNT(*) FROM issue")
	if condition != "" {
		sb.WriteString(" WHERE ")
		sb.WriteString(condition)
		sbCount.WriteString(" WHERE ")
		sbCount.WriteString(condition)
	}
	sb.WriteString(" ORDER BY id DESC")
	sb.WriteString(" LIMIT ")
	sb.WriteString(strconv.Itoa(pageSize))
	sb.WriteString(" OFFSET ")
	sb.WriteString(strconv.Itoa(page * pageSize))
	stmtCount, err := db.Prepare(sbCount.String())
	if err != nil {
		return nil, nil, err
	}
	defer stmtCount.Close()

	var count int
	row := stmtCount.QueryRow(args...)
	if err := row.Scan(&count); err != nil {
		return nil, nil, err
	}

	stmt, err := db.Prepare(sb.String())
	if err != nil {
		return nil, nil, err
	}
	defer stmt.Close()

	var issues []*Issue
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, nil, err
	}
	for rows.Next() {
		issue := new(Issue)
		if err := rows.Scan(&issue.ID, &issue.Ticket, &issue.Result, &issue.Schedule, &issue.BlockTime, &issue.BlockNumber, &issue.Dateline, &issue.DateStr, &issue.BlockHash); err != nil {
			return nil, nil, err
		}
		issues = append(issues, issue)
	}

	pagination := NewPagination(page, count, pageSize)

	return issues, pagination, nil
}
