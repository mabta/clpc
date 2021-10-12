package db

import (
	"testing"

	"github.com/mabta/clpc/internal/cfg"
)

func testInit(t *testing.T) {
	if err := cfg.InitFrom("../../config.json"); err != nil {
		t.Fatal(err)
	}
	if err := Init(cfg.Settings.DB.DSN); err != nil {
		t.Fatal(err)
	}
}

func TestFindIssue(t *testing.T) {
	testInit(t)
	var id uint64 = 211
	issue, err := FindIssue(id)
	if err != nil {
		t.Fatal(err)
	}
	if issue.Ticket != "pk10-5" {
		t.Error("期望 pk10-5，实际", issue.Ticket)
	}
	if issue.BlockNumber != 9018891 {
		t.Error("期望 9018891，实际", issue.BlockNumber)
	}
}
func TestFindIssueBy(t *testing.T) {
	testInit(t)
	condition := "ticket=$1 AND block_number=$2"
	args := []interface{}{"pk10-5", 9018891}
	issue, err := FindIssueBy(condition, args...)
	if err != nil {
		t.Fatal(err)
	}
	if issue.ID != 211 {
		t.Error("期望 211，实际", issue.ID)
	}
}
func TestFindIssueByEmpty(t *testing.T) {
	testInit(t)
	condition := "ticket=$1 AND block_number=$2 AND id=$3"
	args := []interface{}{"not-exists-ticket", 123, 456}
	issue, err := FindIssueBy(condition, args...)
	if !IsNoRows(err) {
		t.Error("期望 sql.ErrNoRows， 实际", err)
	}
	if issue != nil {
		t.Error("期望 nil，实际", issue)
	}
}
func TestSelectIssues(t *testing.T) {
	testInit(t)
	totalPage := 1
	for i := 0; i < totalPage; i++ {
		issues, pagination, err := SelectIssues(i, 30, "")
		if err != nil {
			t.Fatal(err)
		}
		totalPage = pagination.TotalPages
		t.Log(pagination, pagination.PageListToView())
		for _, issue := range issues {
			t.Log(issue)
		}
	}
}
