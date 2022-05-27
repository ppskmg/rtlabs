package apierror

import "encoding/json"

type apiError interface {
	Error() string
	DeveloperError() string
	CodeError() string
}

const (
	ErrorStore         = "Store error"
	ErrorUnknown       = "Unknown error"
	ErrorDecodeString  = "Error DecodeString"
	ErrorTokenValidate = "Error token validate"
)

var (
	//ErrNoRows:         "sql: no rows in result set",

	ErrNoSessionID = New(
		nil, "no session id",
		"No session id",
		"BP-000001")
	ErrNoRows = New(
		nil, "sql: no rows in result set",
		"sql: no rows in result set",
		"BP-000001")
	ErrRecordNotFound = New(
		nil, "record not found",
		"sql: no rows in result set",
		"BP-000001")
	ErrTokenInRedisAccessBlackList = New(
		nil, "token in black list, access denied",
		"token in black list, access denied",
		"BP-000001")
	ErrTokenInRedisRefreshBlackList = New(
		nil, "token in black list, access denied",
		"token in black list, access denied",
		"BP-000001")
	ErrSaveTokenInRedisAccessBlackList = New(
		nil, "Error save token access in redis black list",
		"Error save token access in redis black list",
		"BP-000001")
	ErrDBFindAccessTokenInMainBlackList = New(
		nil, "System error db: err find access token in main black list",
		"System error db: err find access token in main black list",
		"BP-000001")
	ErrCheckMainStorageAccessBlackList = New(
		nil, "Error check main storage access black list",
		"Error check main storage access black list",
		"BP-000001")
	ErrCheckTokenInRedisAccessBlackList = New(
		nil, "Error check token black list",
		"Error check token access in redis black list",
		"BP-000001")
	ErrTokenInRedisOkInMainNotOK = New(
		nil, "Error token",
		"Access токен есть в redis, но нет в main store",
		"BP-000001")
	ErrNoRefreshInMain = New(
		nil, "Error refresh token",
		"Refresh токен отстутсвует в main, но есть access токен (аномалия). Проверка на истёкший access и попытка получения Refresh перед выпуском нового токена.",
		"BP-000001")
	ErrUnmarshalDecodePayload = New(
		nil, "Unmarshal error",
		"Error unmarshal decode payload",
		"BP-000001")
	ErrDeleteExpiredRefresh = New(
		nil, "Record not found",
		"Error delete expired refresh",
		"BP-000001")
	ErrRefreshTokenExpired = New(
		nil, "Refresh token expired",
		"Refresh token expired",
		"BP-000001")
	ErrTokenDoNotMath = New(
		nil, "Token do Not Match",
		"Access token ID не совпадает с ID Access токена в Refresh",
		"BP-000001")
	ErrAddAccessAndRefreshTokenInBlackList = New(
		nil,
		"Store error: Error add access token and refresh token in black list",
		"Access token ID не совпадает с ID Access токена в Refresh, ошибка добавления в Black List",
		"BPS-000000")
	ErrBadRequest = New(
		nil, "Bad Request",
		"Bad request",
		"BP-000001")
	ErrUnauthorized = New(
		nil, "Unauthorized",
		"Complex error for Unauthorized",
		"BP-000001")
	ErrUserAgent = New(
		nil, "Unauthorized",
		"User agent cannot be blank",
		"BP-000002")
	ErrIncorrectEmailOrPassword = New(
		nil, "Incorrect email or password",
		"incorrect email or password",
		"BP-000011")
	ErrIncorrectHeaders = New(
		nil, "incorrect headers",
		"incorrect headers",
		"BP-000003")
	ErrIncorrectFingerprint = New(
		nil, "incorrect fingerprint",
		"incorrect fingerprint",
		"BP-000003")
	ErrSaveFingerPrintInPostgres = New(
		nil, "error save fingerPrint In Postgres",
		"error save fingerPrint In Postgres",
		"BP-000003")
	ErrSaveAccessInRedis = New(
		nil, "redis save token error",
		"redis save token error",
		"BP-000003")
	ErrGetAccessInRedis = New(
		nil, "error get access token in redis",
		"error get access token in redis",
		"BP-000003")
	ErrTokenInvalid = New(
		nil, "incorrect signature",
		"incorrect signature",
		"BP-000003")

	ErrTokenId = New(
		nil, "TokenId different",
		"TokenId different",
		"BP-000003")
	ErrTokenPayload = New(
		nil, "incorrect access token",
		"Token structure: xxx.XXX.xxx — ошибка извлечения полезной нагрузки из access токена полученого после генерации нового токена",
		"BP-000003")
	ErrTokenPayloadMain = New(
		nil, "incorrect access token",
		"Token structure: XXX.xxx — ошибка извлечения полезной нагрузки из access токена полученого из main store",
		"BP-000003")
	ErrTokenSecret = New(
		nil, "incorrect access token",
		"Token structure: xxx.xxx.XXX — ошибка извлечения секретной части токена",
		"BP-000003")
)

type AppError struct {
	Err              error  `json:"err"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developer_message"`
	Code             string `json:"code"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func New(err error, message, developerMessage, code string) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func systemError(err error) *AppError {
	return New(err, "internal system error", err.Error(), "BP-00000")
}
