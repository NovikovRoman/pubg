package pubg

type Error interface {
	GetDetail() string
}

// ErrBadRequest structure.
type ErrBadRequest struct {
	error
	title  string
	detail string
}

// Returns a details.
func (e ErrBadRequest) GetDetail() string {
	return e.detail
}

// Returns a error text.
func (e ErrBadRequest) Error() string {
	return "400 " + e.title
}

// ErrUnauthorized structure.
type ErrUnauthorized struct {
	error
	title  string
	detail string
}

// Returns a details.
func (e ErrUnauthorized) GetDetail() string {
	return e.detail
}

// Returns a error text.
func (e ErrUnauthorized) Error() string {
	return "401 " + e.title
}

// ErrNotFound structure.
type ErrNotFound struct {
	error
	title  string
	detail string
}

// Returns a details.
func (e ErrNotFound) GetDetail() string {
	return e.detail
}

// Returns a error text.
func (e ErrNotFound) Error() string {
	return "404 " + e.title
}

// ErrUnsupportedMediaType structure.
type ErrUnsupportedMediaType struct {
	error
	title  string
	detail string
}

// Returns a details.
func (e ErrUnsupportedMediaType) GetDetail() string {
	return e.detail
}

// Returns a error text.
func (e ErrUnsupportedMediaType) Error() string {
	return "415 " + e.title
}

// ErrTooManyRequest structure.
type ErrTooManyRequest struct {
	error
	title  string
	detail string
}

// Returns a details.
func (e ErrTooManyRequest) GetDetail() string {
	return e.detail
}

// Returns a error text.
func (e ErrTooManyRequest) Error() string {
	return "429 Too many request"
}
