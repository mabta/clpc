package handler

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mabta/clpc/defs"
	"github.com/mabta/clpc/defs/draw"
	"github.com/mabta/clpc/internal/db"
	"github.com/mabta/clpc/internal/lottery"
	"github.com/mabta/clpc/internal/redis"
)

func DefaultHandler(describies []*lottery.Describe, block *defs.Block) {
	defer func() {
		if p := recover(); p != nil {
			log.Println("事情大条了：", p)
		}
	}()
	for _, d := range describies {
		d.SetBlock(block)
		for _, schedule := range d.Schedule() {
			// 判断区块时间
			if !schedule.IsIn(block.Time) {
				log.Println("不在开奖时间内，跳过")
				continue
			}
			// 是否已开奖
			if isDrawed(d.Name, schedule.Start) {
				log.Println("已开奖，跳过")
				continue
			}
			// 是否第一块
			if hasFirstBlock(d.Name, schedule.Start) {
				// 第2块
				// 开奖
				result := d.Draw()
				// 保存开奖结果
				issue := db.NewIssue(d.Name, drawResult2Str(result), schedule.IssueStr(), block.Time, block.Number)
				id, err := db.InsertIssue(issue)
				if err != nil {
					panic(err)
				}
				log.Println("开奖结果：", result, id)
			} else {
				log.Println("第一块，保存到暂存区")
				cacheFirstBlock(d.Name, schedule.Start, block.Number, d.Duration)
			}
		}
	}
}

func redisKeyIsExists(key string) bool {
	exists, err := redis.Exists(key)
	if err != nil {
		panic("redis.Exists:" + err.Error())
	}
	return exists
}

func redisSet(key string, value interface{}, duration time.Duration) {
	if err := redis.SetExpiredDuration(key, value, duration); err != nil {
		panic("redis.SetExpiredDuration:" + err.Error())
	}
}

func firstBlockKey(ticketName string, schedule uint64) string {
	return fmt.Sprintf("first-block-%s-%d", ticketName, schedule)
}
func drawedKey(ticketName string, schedule uint64) string {
	return fmt.Sprintf("drawed-%s-%d", ticketName, schedule)
}

func hasFirstBlock(ticketName string, schedule uint64) bool {
	key := firstBlockKey(ticketName, schedule)
	return redisKeyIsExists(key)
}
func isDrawed(ticketName string, schedule uint64) bool {
	key := drawedKey(ticketName, schedule)
	return redisKeyIsExists(key)
}

func cacheFirstBlock(ticketName string, schedule uint64, blockNumber uint64, duration time.Duration) {
	key := firstBlockKey(ticketName, schedule)
	redisSet(key, blockNumber, duration)
}

func saveDrawedBlock(ticketName string, schedule uint64, blockNumber uint64, duration time.Duration) {
	key := drawedKey(ticketName, schedule)
	redisSet(key, blockNumber, duration)
}

func drawResult2Str(result []draw.Result) string {
	sb := strings.Builder{}
	for i, r := range result {
		sb.WriteString(fmt.Sprintf("%d", int(r)))
		if i < len(result)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}
