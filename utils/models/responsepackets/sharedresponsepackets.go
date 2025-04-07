package responsepackets

import "encoding/json"

type CommonResponsePacket struct {
	Success bool            `json:"success"`
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"` // Use json.RawMessage to handle various structures for "data"
}
// ErrorResponse Structure for the first response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
}


//what does omitempty do
// omitempty means that if the field is empty (zero value), it will be omitted from the JSON output.
// This is useful for reducing the size of the JSON response and avoiding sending unnecessary data.
// CommonErrorResponse Structure for the second response

type CommanErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"error_message"`
	Code    int         `json:"error_code"`
	AppCode int         `json:"app_code"`
	LogId   string      `json:"log_id"`
	Data    interface{} `json:"data,omitempty"`   // Optional field to include response data
	Errors  interface{} `json:"errors,omitempty"` // Optional field to include error details
}
