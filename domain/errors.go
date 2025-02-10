package domain

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrInvalidIDNumber = &ApiError{
		Code:    10001,
		Message: "非法的身份证号",
	}
)
