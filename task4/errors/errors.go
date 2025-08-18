package errors

// 自定义错误类型，包含错误信息和对应的HTTP状态码
type AppError struct {
	Code    int    // HTTP状态码
	Message string // 错误信息（给前端展示）
	Err     error  // 原始错误（用于后端日志）
}

// 实现error接口
func (e *AppError) Error() string {
	return e.Message
}

func NewError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
