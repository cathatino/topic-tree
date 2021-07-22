package authmethod

const (
	EmptyNickName = string("")
)

type AuthMethodName string

var (
	NoAuthMethodName = AuthMethodName("no_auth")
)
