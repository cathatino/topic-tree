package authmethod

// NoAuthMethod - a no validation required auth
type NoAuthMethod struct{}

// NoAuthBody defines the required field to be checked by NoAuthMethod.IsAllowed func
type NoAuthBody struct {
	NickName string
}

func (nam NoAuthMethod) IsAllowed(noAuthBody interface{}) error {
	covertedNoAuthBody, ok := noAuthBody.(NoAuthBody)
	if !ok {
		return WrongAuthBodyTypeErr
	}
	if covertedNoAuthBody.NickName == EmptyNickName {
		return NoNickNameErr
	}
	return nil
}

func NewNoAuthMethod() AuthMethod {
	return &NoAuthMethod{}
}
