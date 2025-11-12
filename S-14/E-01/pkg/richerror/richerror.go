package richerror

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
)

type RichError struct {
	operation      string
	wrapperedError error
	message        string
	kind           Kind
	meta           map[string]interface{}
}

func New(err error, operation, message string, kind Kind, meta map[string]interface{}) RichError {
	return RichError{
		operation:      operation,
		wrapperedError: err,
		message:        message,
		kind:           kind,
		meta:           meta,
	}
}