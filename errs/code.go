package errs

const (
	Unknown              = 1000
	DBError              = 1001
	InvalidParam         = 1002
	NotFound             = 1003
	InvalidLotteryFormat = 1201
	BadRequest           = 9999
)

func ToDbError(err error) *Error {
	return New(DBError, "数据库操作失败", err)
}
func NewInvalidParam() *Error {
	return New(InvalidParam, "参数错误", nil)
}
func NotFoundError() *Error {
	return New(NotFound, "未找到满足条件的记录", nil)
}
