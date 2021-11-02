package blocktime

import "time"

// ToTime 区块时间转为 Go 的 time.Time
func ToTime(blockTime uint64) time.Time {
	return time.Unix(int64(blockTime), 0)
}

// DateStr 区块时间转为字符表示的日期
func DateStr(blockTime uint64) string {
	return ToTime(blockTime).Format("20060102")
}

// TimeStr 区块时间转为字符表示的时间
func TimeStr(blockTime uint64) string {
	return ToTime(blockTime).Format("15:04:05")
}

func DateTimeStr(blockTime uint64) string {
	return ToTime(blockTime).Format("2006/01/02 15:04:05")
}
func DateTimeStrToApi(blockTime uint64) string {
	return ToTime(blockTime).Format("2006-01-02 15:04:05")
}
func IssueStr(blockTime uint64) string {
	return ToTime(blockTime).Format("20060102 15:04")
}

// ToDateUnix 区块时间转为日期(零点)对应的时间戳
func ToDateUnix(blocktime uint64) uint64 {
	tt := ToTime(blocktime)
	ntt := time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, time.Local)
	return uint64(ntt.Unix())
}
