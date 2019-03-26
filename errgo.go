package errgo

// ErrorCheck - func to check error
type ErrorCheck func(error)

// ErrorHandler - func that understand how to handle error
type ErrorHandler func(error)

// NewErrorCheck - return error check function
func NewErrorCheck(handler ErrorHandler) ErrorCheck {
	return func(err error) {
		if err != nil {
			handler(err)
		}
	}
}
