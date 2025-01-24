package defs

type ErrCode int32

const (
	ErrCodeOK               ErrCode = 0 // 正常响应
	ErrCodeSystemPanic      ErrCode = 1 // 系统崩溃
	ErrCodeSystemError      ErrCode = 2 // 系统错误
	ErrCodeRequestParamsErr ErrCode = 3 // 请求参数错误
)

type CustomError struct {
	Code ErrCode
	Err  error
}

func (c *CustomError) GetCode() int {
	if c == nil {
		return 0
	}
	return int(c.Code)
}

func (c *CustomError) GetErr() string {
	if c == nil {
		return ""
	}
	return c.Err.Error()
}

func NewCustomError(code ErrCode, err error) *CustomError {
	return &CustomError{
		Code: code,
		Err:  err,
	}
}
