package app

type AppError struct {
	Code    int
	Message string
}

var (
	ErrBadRequest = &AppError{400, "Bad request"}
)

func (e *AppError) Error() string {
	return e.Message
}
