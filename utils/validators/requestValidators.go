package validators

import (
	"fmt"
	"reflect"

	"github.com/TechnoDiktator/ginTutorial/utils/models/requestpackets"
	"github.com/sirupsen/logrus"
)

func ValidateRequestForTempObject(requestPacket requestpackets.RequestPacketTypeOne) (bool, []string, []string, []string, []string) {

	logrus.Info("Validating request: ", requestPacket)

	//maybe we can use reflection to validate the struct
	// Check if the required fields are present

	/*
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

	*/

	// var dbobject *dbmodels.Temp

	val := reflect.ValueOf(requestPacket)
	typ := reflect.TypeOf(requestPacket)

	if val.Kind() != reflect.Struct {
		return false, []string{"provided value is not a struct"}, []string{}, []string{}, []string{}
	}

	expectedTypes := map[string]reflect.Kind{
		"TicketNum":           reflect.String,
		"AoiSiteID":           reflect.Int,
		"GateID":              reflect.Int,
		"DeviceType":          reflect.Slice,
		"RequestTime":         reflect.Int64,
		"TotalAmount":         reflect.Float64,
		"ExitGateOperatorId":  reflect.Int,
		"EntryGateOperatorId": reflect.Int,
		"EntryGate":           reflect.Int,
		"ExitGate":            reflect.Int,
		"EntryTime":           reflect.String,
		"ExitTime":            reflect.String,
		"DontUpdateTicket":    reflect.Bool,
		"EntryEpochTime":      reflect.Int64,
		"ExitEpochTime":       reflect.Int64,
	}

	var typeErrors, emptyErrors, fetchErrors, dataMatchingError []string

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		//fieldType := typ.Field(i)
		fieldName := field.Name
		fieldValue := val.Field(i)

		// fieldTypeName := fieldType.Type.Name()
		// // fieldTag := fieldType.Tag.Get("json")
		// // fieldValue := field.Interface()
		expectedType, ok := expectedTypes[fieldName]
		if !ok {
			typeErrors = append(typeErrors, fmt.Sprintf("field %s: unexpected field", fieldName))
			continue
		}

		// Validate the field for type and empty error
		err, errType := validateField(fieldName, fieldValue, expectedType)
		if err != "" {
			if errType == "type error" {
				typeErrors = append(typeErrors, err)
			} else if errType == "empty error" {
				emptyErrors = append(emptyErrors, err)
			}
		}
	}
	return true, typeErrors, emptyErrors, fetchErrors, dataMatchingError
}

func validateField(fieldName string, fieldValue reflect.Value, expectedType reflect.Kind) (string, string) {
	/*
		// validateField checks if a field's type matches the expected type and whether the field's value is empty.
		// Parameters:
		// - fieldName: The name of the field being validated.
		// - fieldValue: The reflect.Value representing the field's value.
		// - expectedType: The expected reflect.Kind type of the field.
		// Returns:
		// - A string describing the error, if any.
		// - A string representing the type of error ("type error" or "empty error").
	*/
	valType := fieldValue.Kind()

	// Check if the field's type matches the expected type
	if valType != expectedType {
		return fmt.Sprintf("field %s: expected type %s, got %s", fieldName, expectedType, valType), "type error"
	}

	// Check if the field's value is empty
	if fieldValue.IsZero() {
		return fmt.Sprintf("field %s is empty", fieldName), "empty error"
	}

	return "", ""
}
