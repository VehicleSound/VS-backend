package usecase

type ErrNotFound struct {
	Err error
}

func NewErrNotFound(err error) ErrNotFound {
	return ErrNotFound{Err: err}
}
