package errors

type ErrProductNotFound struct{}

func (e ErrProductNotFound) Error() string {
	return "Product not found"
}
