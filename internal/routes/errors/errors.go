package errors

import "fmt"

const (
	errorType = "error"

	NotValidToken        = "NotValidToken"
	notValidTokenMessage = "token isn't valid"

	Internal        = "Internal"
	internalMessage = "internal"

	DecodeRequest        = "DecodeRequest"
	decodeRequestMessage = "request wasn't decoded"

	MapRequest        = "MapRequest"
	mapRequestMessage = "request wasn't mapped to domain model"

	Unknown        = "Unknown"
	unknownMessage = "unknown"
)

type Error struct {
	Type    string `json:"type"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewError(code string, err error) error {
	var message string

	switch code {
	case NotValidToken:
		message = notValidTokenMessage
	case Internal:
		message = internalMessage
	case DecodeRequest:
		message = decodeRequestMessage
	case MapRequest:
		message = mapRequestMessage
	default:
		message = unknownMessage
	}

	if err != nil {
		message = fmt.Sprintf("%s (%s)", message, err)
	}

	return &Error{
		Type:    errorType,
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	return e.Message
}
