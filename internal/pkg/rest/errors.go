package rest

type AuthenticationError struct{}

func (e AuthenticationError) Error() string {
	return "Wrong email or password"
}
