package apperror

var MsgFlags = map[error]int32{
	ErrorInvalidUsername: 10000,
	ErrorInvalidPassword: 10001,
	ErrorInvalidFullname: 10002,
	ErrorEmptyField:      10003,
	ErrorSpaceDetected:   10004,
	ErrorWrongDateFormat: 10005,

	ErrorEntryExists:             20000,
	ErrorInputInvalid:            20001,
	ErrorCookieNotFound:          20002,
	ErrorCookieOutdated:          20003,
	ErrorBlockedIP:               20004,
	ErrorNotAuthenticated:        20005,
	ErrorNotAuthorized:           20006,
	ErrorJWTAuthenticationFailed: 20007,
	ErrorCannotDeleteThisEntity:  20008,
	ErrorUsernameAlreadyExist:    20009,
	ErrorNameAlreadyExist:        20010,
	ErrorPasswordIncorrect:       20011,
	ErrorUserNotFound:            20012,

	ErrorInternal:         30000,
	ErrorEntryNotExist:    30001,
	ErrorConflict:         30002,
	ErrorBindJSON:         30003,
	ErrorContextLostValue: 30004,
	ErrorNotSupportedYet:  30005,
}

func GetCode(e error) int32 {
	code, ok := MsgFlags[e]
	if ok {
		return code
	}

	return 0
}
