package authmethod

// AuthMethod intefrace for auth validation purpose
type AuthMethod interface {
	IsAllowed(authBody interface{}) error
}

// GetAuthMethodByAuthMethodName fetchs AuthMethod by authMethodName
func GetAuthMethodByAuthMethodName(
	authMethodName AuthMethodName,
) (authMehtod AuthMethod, err error) {
	if authMethodName == NoAuthMethodName {
		return NewNoAuthMethod(), nil
	}
	return nil, InvalidAuthMethodName
}
