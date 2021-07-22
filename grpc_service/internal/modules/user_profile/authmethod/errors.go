package authmethod

import "errors"

type AuthMehthodError error

var (
	// Auth Method Related errors
	InvalidAuthMethodName = AuthMehthodError(errors.New("invalid_auth_method_name"))
	NoNickNameErr         = AuthMehthodError(errors.New("no_nick_name_error"))
	WrongAuthBodyTypeErr  = AuthMehthodError(errors.New("wrong_auth_body_type"))
)
