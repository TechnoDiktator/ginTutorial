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
