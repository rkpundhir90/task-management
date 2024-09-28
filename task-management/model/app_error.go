package model

var emptyAppError = AppError{}

type AppError struct {
	Error     error
	Message   string
	Code      int
	ErrorCode int `json:"errorCode,omitempty"`
}

func (ae AppError) IsEmpty() bool {
	return ae == emptyAppError
}

func (ae AppError) IsNotEmpty() bool {
	return ae != emptyAppError
}

func NoError() AppError {
	return emptyAppError
}

type Error struct {
	Error   string
	Message string
}
