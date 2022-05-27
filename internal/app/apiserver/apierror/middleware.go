package apierror

type ErrorsMiddleware struct {
}
type Errors struct {
	ErrUnauthorized   *AppError
	ErrUserAgentBlank *AppError
	ErrNotFound       *AppError
}
