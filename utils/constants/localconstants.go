package constants
type State struct {
	Success         bool
	Message         string
	ErrorCode       int
	HttpsStatusCode int
}

var (


	StateInvalidRequestPacket = State{
		Success:         false,
		Message:         "Invalid Request packet format.",
		ErrorCode:       2116,
		HttpsStatusCode: 400,
	}

	StateValidationFailed = State{
		Success:         false,
		Message:         "Validation failed. Please refer the documentation.",
		ErrorCode:       2117,
		HttpsStatusCode: 400,
	}

	StateUnknown = State{
		Success:         false,
		Message:         "Unwanted response received from Venodr",
		ErrorCode:       2102,
		HttpsStatusCode: 400,
	}

)
