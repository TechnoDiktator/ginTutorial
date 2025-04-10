package requestpackets



//Example of request packet
type RequestPacketTypeOne struct {
	TicketNum           string   `json:"ticket_num"`
	AoiSiteID           int      `json:"aoi_site_id"`
	GateID              int      `json:"gate_id"`
	DeviceType          []string `json:"device_type"`
	RequestTime         int64    `json:"request_time"`
	TotalAmount         float64  `json:"total_amount"`
	ExitTime            string   `json:"exit_time"`
	EntryTime           string   `json:"entry_time"`
	ExitGate            int      `json:"exit_gate_id"`
	EntryGate           int      `json:"entry_gate_id"`
	EntryGateOperatorId int      `json:"entry_gate_operator_id"`
	ExitGateOperatorId  int      `json:"exit_gate_operator_id"`
	DontUpdateTicket    bool     `json:"without_saving_ticket"`
	EntryEpochTime      int64    `json:"entry_time_epoch"`
	ExitEpochTime       int64    `json:"exit_time_epoch"`


}


type CreateSurveyRequest struct {
	Heading     string   `json:"heading" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Questions   []string `json:"questions" binding:"required"`
}


