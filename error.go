package imageserver

// Error displayable to the end user.
//
// It is used in the http server: the errors of this type are displayed, the others are shown as "internal error".
type Error struct {
	Text string
}

// NewError creates a new Error
func NewError(text string) *Error {
	return &Error{
		Text: text,
	}
}

func (err *Error) Error() string {
	return err.Text
}
