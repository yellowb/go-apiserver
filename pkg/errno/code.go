package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// Create user errors
	EncryptErr = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	UserFieldMissed = &Errno{Code: 20102, Message: "The username was missed."}
	PwdFieldMissed = &Errno{Code: 20103, Message: "The password was missed."}
	TokenInvalid = &Errno{Code: 20104, Message: "The token was invalid."}
)
