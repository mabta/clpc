package errs

type Error struct {
	Code  int
	Msg   string
	Cause error
}

func (e *Error) Error() string {
	return e.Msg
}

func New(code int, msg string, cause error) *Error {
	return &Error{
		Code:  code,
		Msg:   msg,
		Cause: cause,
	}
}
