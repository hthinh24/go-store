package errors

type ErrPasswordMismatch struct{}

func (p ErrPasswordMismatch) Error() string {
	return "Password & confirm password not match"
}
