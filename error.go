package pubg

type Error interface {
	GetDetail() string
}

// ErrBadRequest --------------------------------------------------------------------------------
type ErrBadRequest struct {
	error
	title  string
	detail string
}

func (e ErrBadRequest) GetDetail() string {
	return e.detail
}

func (e ErrBadRequest) Error() string {
	return "400 " + e.title
}

// ErrUnauthorized --------------------------------------------------------------------------------
type ErrUnauthorized struct {
	error
	title  string
	detail string
}

func (e ErrUnauthorized) GetDetail() string {
	return e.detail
}

func (e ErrUnauthorized) Error() string {
	return "401 " + e.title
}

// ErrNotFound --------------------------------------------------------------------------------
type ErrNotFound struct {
	error
	title  string
	detail string
}

func (e ErrNotFound) GetDetail() string {
	return e.detail
}

func (e ErrNotFound) Error() string {
	return "404 " + e.title
}

// ErrUnsupportedMediaType --------------------------------------------------------------------------------
type ErrUnsupportedMediaType struct {
	error
	title  string
	detail string
}

func (e ErrUnsupportedMediaType) GetDetail() string {
	return e.detail
}

func (e ErrUnsupportedMediaType) Error() string {
	return "415 " + e.title
}

// ErrTooManyRequest --------------------------------------------------------------------------------
type ErrTooManyRequest struct {
	error
	title  string
	detail string
}

func (e ErrTooManyRequest) GetDetail() string {
	return e.detail
}

func (e ErrTooManyRequest) Error() string {
	return "429 Too many request"
}
