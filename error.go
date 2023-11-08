package pubg

// Error interface service API errors
type Error interface {
	GetDetail() string
}

// ErrBadRequest structure.
type ErrBadRequest struct {
	title  string
	detail string
}

// GetDetail returns a details.
func (e ErrBadRequest) GetDetail() string {
	return e.detail
}

// Error returns a error text.
func (e ErrBadRequest) Error() string {
	return "400 " + e.title
}

// ErrUnauthorized structure.
type ErrUnauthorized struct {
	title  string
	detail string
}

// GetDetail returns a details.
func (e ErrUnauthorized) GetDetail() string {
	return e.detail
}

// Error returns a error text.
func (e ErrUnauthorized) Error() string {
	return "401 " + e.title
}

// ErrNotFound structure.
type ErrNotFound struct {
	title  string
	detail string
}

// GetDetail returns a details.
func (e ErrNotFound) GetDetail() string {
	return e.detail
}

// Error returns a error text.
func (e ErrNotFound) Error() string {
	return "404 " + e.title
}

// ErrUnsupportedMediaType structure.
type ErrUnsupportedMediaType struct {
	title  string
	detail string
}

// GetDetail returns a details.
func (e ErrUnsupportedMediaType) GetDetail() string {
	return e.detail
}

// Error returns a error text.
func (e ErrUnsupportedMediaType) Error() string {
	return "415 " + e.title
}

// ErrTooManyRequest structure.
type ErrTooManyRequest struct {
	detail string
}

// GetDetail returns a details.
func (e ErrTooManyRequest) GetDetail() string {
	return e.detail
}

// Error returns a error text.
func (e ErrTooManyRequest) Error() string {
	return "429 Too many request"
}
