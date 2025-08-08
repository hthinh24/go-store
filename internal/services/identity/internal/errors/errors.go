package errors

type ErrUserNotFound struct{}

func (e ErrUserNotFound) Error() string {
	return "User not found"
}

type ErrUserNotActive struct{}

func (e ErrUserNotActive) Error() string {
	return "User is not active"
}

type ErrPasswordMismatch struct{}

func (p ErrPasswordMismatch) Error() string {
	return "Password & confirm password not match"
}
