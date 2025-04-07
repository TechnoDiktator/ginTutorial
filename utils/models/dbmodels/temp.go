package dbmodels

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/lib/pq"
)







































































type Temp struct {
	ID                             int             `json:"id" db:"id" gorm:"primary_key"`
	Type                           string          `json:"type" db:"type"`
	TicketNum                      string          `json:"ticket_num" db:"ticket_num"`
	TicketState                    string          `json:"ticket_state" db:"ticket_state"`
	UploadStatus                   bool            `json:"upload_status" db:"upload_status"`
	IntimationStatus               bool            `json:"intimation_status" db:"intimation_status"`
	VehiclePlateNum                sql.NullString  `json:"vehicle_plate_num" db:"vehicle_plate_num"`
	EntryGateOperator              sql.NullInt16   `json:"entry_gate_operator" db:"entry_gate_operator"`
	ExitGateOperator               sql.NullInt16   `json:"exit_gate_operator" db:"exit_gate_operator"`
	EntryGate                      sql.NullInt16   `json:"entry_gate" db:"entry_gate"`
	EntryGateName                  sql.NullString  `json:"entry_gate_name" db:"entry_gate_name"`
	ExitGate                       sql.NullInt16   `json:"exit_gate" db:"exit_gate"`
	EntryTime                      *time.Time      `json:"entry_time" db:"entry_time"`
	ExitTime                       *time.Time      `json:"exit_time" db:"exit_time"`
	AOISite                        sql.NullInt16   `json:"aoi_site" db:"aoi_site"`
	WinnerOCR                      sql.NullString  `json:"winner_ocr" db:"winner_ocr"`
	ExitWinnerOCR                  sql.NullString  `json:"exit_winner_ocr" db:"exit_winner_ocr"`
	WinnerTagID                    sql.NullString  `json:"winner_tag_id" db:"winner_tag_id"`
	ExitWinnerTagID                sql.NullString  `json:"exit_winner_tag_id" db:"exit_winner_tag_id"`
	EntryANPRList                  pq.StringArray  `json:"entry_anpr_list" db:"entry_anpr_list"`
	ExitANPRList                   pq.StringArray  `json:"exit_anpr_list" db:"exit_anpr_list"`
	TicketOpeningMethod            sql.NullString  `json:"ticket_opening_method" db:"ticket_opening_method"`
	TicketClosingMethod            sql.NullString  `json:"ticket_closing_method" db:"ticket_closing_method"`
	ExitDetectedBrand              sql.NullString  `json:"exit_detected_brand" db:"exit_detected_brand"`
	EntryCorrectedData             sql.NullString  `json:"entry_corrected_data" db:"entry_corrected_data"`
	ExitCorrectedData              sql.NullString  `json:"exit_corrected_data" db:"exit_corrected_data"`
	VehicleType                    sql.NullString  `json:"vehicle_type" db:"vehicle_type"`
	TransactionID                  sql.NullString  `json:"transaction_id" db:"transaction_id"`
	EntryNumber                    sql.NullInt16   `json:"entry_number" db:"entry_number"`
	ReservedCorporateVehicleTicket sql.NullBool    `json:"reserved_corporate_vehicle_ticket" db:"reserved_corporate_vehicle_ticket"`
	ReservedCorporate              sql.NullInt16   `json:"reserved_corporate" db:"reserved_corporate"`
	SitePassID                     sql.NullString  `json:"site_pass_id" db:"site_pass_id"`
	SitePassANPR                   sql.NullString  `json:"site_pass_anpr" db:"site_pass_anpr"`
	TotalAmount                    sql.NullFloat64 `json:"total_amount" db:"total_amount"`
	DiscountAmount                 sql.NullFloat64 `json:"discount_amount" db:"discount_amount"`
	Agent                          sql.NullString  `json:"agent" db:"agent"`
	Status                         sql.NullString  `json:"status" db:"status"`
	IsFOC                          sql.NullBool    `json:"is_foc" db:"is_foc"`
	EntryRFIDData                  json.RawMessage `json:"entry_rfid_data" db:"entry_rfid_data"`
	ExitRFIDData                   json.RawMessage `json:"exit_rfid_data" db:"exit_rfid_data"`
	EntryEventRFIDData             json.RawMessage `json:"entry_event_rfid_data" db:"entry_event_rfid_data"`
	QRCode                         sql.NullString  `json:"qr_code" db:"qr_code"`
	FastagErrorCodes               sql.NullString  `json:"fastag_error_codes" db:"fastag_error_codes"`
	CreatedOn                      *time.Time      `json:"created_on" db:"created_on"`
	TicketClosureStartTime         *time.Time      `json:"ticket_closure_start_time" db:"ticket_closure_start_time"`
	TicketClosureEndTime           *time.Time      `json:"ticket_closure_end_time" db:"ticket_closure_end_time"`
	BarrierOpeningTime             *time.Time      `json:"barrier_opening_time" db:"barrier_opening_time"`
	UpdatedOn                      *time.Time      `json:"updated_on" db:"updated_on"`
	GatesFailedToUpload            pq.StringArray  `json:"gates_failed_to_upload" db:"gates_failed_to_upload"`
	UploadToCloudStatus            sql.NullString  `json:"upload_to_cloud_status" db:"upload_to_cloud_status"`
	UploadToGatesStatus            sql.NullString  `json:"upload_to_gates_status" db:"upload_to_gates_status"`
	FOCReasonForCancelling         sql.NullString  `json:"foc_reason_for_cancelling" db:"foc_reason_for_cancelling"`
	ReasonCodeForCancelling        sql.NullString  `json:"reason_code_for_cancelling" db:"reason_code_for_cancelling"`
	JetsonSelfIntimationStatus     sql.NullBool    `json:"jetson_self_intimation_status" db:"jetson_self_intimation_status"`
	PlateColorDetected             sql.NullString  `json:"plate_color_detected" db:"plate_color_detected"`
	PriceCalculationBasedOn        sql.NullString  `json:"price_calculation_based" db:"price_calculation_based"`
	ListOfAdditionalReasons        sql.NullString  `json:"list_of_additional_reasons" db:"list_of_additional_reasons"`
	PaymentReferenceTimestamp      sql.NullString  `json:"payment_reference_timestamp" db:"payment_reference_timestamp"`
	TransactionHistory             json.RawMessage `json:"transaction_history" db:"transaction_history"`
	TicketClosedVia                sql.NullString  `json:"ticket_closed_via" db:"ticket_closed_via"`
	CouponUsedID                   sql.NullInt16   `json:"coupon_used_id" db:"coupon_used_id"`
	EventID                        sql.NullInt64   `json:"event_id" db:"event_id"`
	PriceCaculationBasedOn         sql.NullString  `json:"price_caculation_based_on" db:"price_caculation_based_on"`
	TicketCreatedClosedAtExit      sql.NullBool    `json:"ticket_created_closed_at_exit" db:"ticket_created_closed_at_exit"`
}
