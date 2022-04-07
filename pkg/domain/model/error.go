package model

var (
	ErrIDRequired      = NewError("id.required")
	ErrIDInvalid       = NewError("id.invalid_value")
	ErrUserIDRequired  = NewError("user_id.required")
	ErrContentRequired = NewError("content.required")
	ErrContentInvalid  = NewError("content.invalid_value")

	ErrCreatedFromInvalid = NewError("created_from.invalid")
	ErrCreatedToInvalid   = NewError("created_to.invalid")

	ErrSizeFromInvalid = NewError("size_from.invalid")
	ErrSizeToInvalid   = NewError("size_to.invalid")

	ErrImageNotFound = NewError("image.not_found")
)

type Error struct {
	text string
}

func NewError(text string) Error {
	return Error{
		text: text,
	}
}

func (e Error) Error() string {
	return e.text
}
