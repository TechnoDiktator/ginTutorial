package pg

import (
	"context"

	"time"

	localmodels "github.com/TechnoDiktator/ginTutorial/utils/models/dbmodels"
	"github.com/sirupsen/logrus"
)

func (pg *LocalPgClient) GetLastTicketObject(defaultGate int) (*localmodels.Temp, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var lastTicketObject localmodels.Temp
	query := `SELECT * FROM ticket_event_ticket WHERE (entry_gate = $1 OR exit_gate = $1) ORDER BY created_on DESC LIMIT 1`
	res, err := pg.DB.QueryxContext(ctx, query, defaultGate)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	if res.Next() {
		err = res.StructScan(&lastTicketObject)
		if err != nil {
			logrus.Error("Error while fetching latest ticket object", err)
			return nil, err
		}
	} else {
		logrus.Infof("No ticket object found for gate %d", defaultGate)
		return nil, nil
	}
	logrus.Infof("Latest Ticket Fetched %s ", lastTicketObject.TicketNum)

	return &lastTicketObject, nil
}





































































































/*


func (pg *LocalPgClient) GetLastTicketClosedSameAsCurrent() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	log.Infof("WE ARE INTO TICKET FETHCING FUNC")
	var ticketObj localdqrdbmodels.TicketEventTicket
	query := `SELECT * FROM ticket_event_ticket WHERE agent = 'cash' ORDER BY id DESC LIMIT 1`
	// Execute the query with parameterized input
	log.Infof("QEUERY", query)

	res, err := pg.DB.QueryxContext(ctx, query)
	if err != nil {
		log.Error("failed to execute query: %w", err)
		return "", fmt.Errorf("failed to execute query: %w", err)
	}
	defer res.Close()

	log.Infof("QUERY RESP", res)

	// Iterate over result rows
	if res.Next() {
		err = res.StructScan(&ticketObj)
		if err != nil {
			log.Error("failed to scan ticket object: %w", err)
			return "", fmt.Errorf("failed to scan ticket object: %w", err)
		}
	} else {
		// No records found
		log.Info("no object found with the given ticket number")
		return "", errors.New("no object found with the given ticket number")
	}

	// Handle any error encountered while iterating over results
	if err = res.Err(); err != nil {
		log.Error("error while iterating over results: %w", err)
		return "", fmt.Errorf("error while iterating over results: %w", err)
	}
	// log.Infof("WE GOT THE TICKET OBJ", &ticketObj)
	return ticketObj.TicketNum, nil
}

func (pg *LocalPgClient) GetCashTicketObjectByTicketNum(ticketNum string) (*localdqrdbmodels.TicketEventTicket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	log.Infof("WE ARE INTO TICKET FETHCING FUNC")
	var ticketObj localdqrdbmodels.TicketEventTicket
	query := `SELECT * FROM ticket_event_ticket WHERE ticket_num=$1 AND agent='cash'` // Fixed the query here
	// Execute the query with parameterized input
	log.Infof("QUERY: %s", query)

	res, err := pg.DB.QueryxContext(ctx, query, ticketNum)
	if err != nil {
		log.Error("failed to execute query: %w", err)
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer res.Close()

	log.Infof("QUERY RESP: %v", res)

	// Iterate over result rows
	if res.Next() {
		err = res.StructScan(&ticketObj)
		if err != nil {
			log.Error("failed to scan ticket object: %w", err)
			return nil, fmt.Errorf("failed to scan ticket object: %w", err)
		}
	} else {
		// No records found
		log.Info("no object found with the given ticket number")
		return nil, errors.New("no object found with the given ticket number")
	}

	// Handle any error encountered while iterating over results
	if err = res.Err(); err != nil {
		log.Error("error while iterating over results: %w", err)
		return nil, fmt.Errorf("error while iterating over results: %w", err)
	}
	// log.Infof("WE GOT THE TICKET OBJ: %v", &ticketObj)
	log.Infof("QUERY RESP: %v", &ticketObj)
	return &ticketObj, nil
}

func (pg *LocalPgClient) GetTicketObjectByTicketNum(ticketNum string) (*localdqrdbmodels.TicketEventTicket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// ticketObj := &Ticket{}
	log.Infof("WE ARE INTO TICKET FETHCING FUNC")
	var ticketObj localdqrdbmodels.TicketEventTicket
	query := `SELECT * FROM ticket_event_ticket WHERE ticket_num=$1 ORDER BY id DESC
		LIMIT 1`
	// Execute the query with parameterized input
	//log.Infof("QEUERY", query)

	res, err := pg.DB.QueryxContext(ctx, query, ticketNum)
	if err != nil {
		log.Error("failed to execute query: %w", err)
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer res.Close()

	//log.Infof("QUERY RESP", res)

	// Iterate over result rows
	if res.Next() {
		err = res.StructScan(&ticketObj)
		if err != nil {
			log.Error("failed to scan ticket object: %w", err)
			return nil, fmt.Errorf("failed to scan ticket object: %w", err)
		}
	} else {
		// No records found
		log.Info("no object found with the given ticket number")
		return nil, nil
	}

	// Handle any error encountered while iterating over results
	if err = res.Err(); err != nil {
		log.Error("error while iterating over results: %w", err)
		return nil, fmt.Errorf("error while iterating over results: %w", err)
	}
	// log.Infof("WE GOT THE TICKET OBJ", &ticketObj)
	return &ticketObj, nil
}

func (pg *LocalPgClient) GetTicketObjectByDqrTransactionId(dqrTransactionId string) (*localdqrdbmodels.TicketEventTicket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var ticketObj localdqrdbmodels.TicketEventTicket
	query := `SELECT * FROM ticket_event_ticket WHERE dqr_transaction_id=$1`

	// Execute the query with parameterized input
	res, err := pg.DB.QueryxContext(ctx, query, dqrTransactionId)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	// Iterate over result rows
	if res.Next() {
		err = res.StructScan(&ticketObj)
		if err != nil {
			return nil, err
		}
	} else {
		// No records found
		return nil, errors.New("no object found with the given value")
	}

	// Handle any error encountered while iterating over results
	if err = res.Err(); err != nil {
		return nil, err
	}

	return &ticketObj, nil
}

/*
implementation for field map
fields := map[string]interface{}{
    "name":   "Updated Name",
    "field2": "Updated Field2 Value",
}
*/

// func (pg *LocalPgClient) UpdateTicketObject(ticketNum string, fields map[string]interface{}) (bool, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	if len(fields) == 0 {
// 		return false, errors.New("no fields provided for updating ticket object")
// 	}

// 	// Build the SET clause dynamically
// 	setClauses := make([]string, 0, len(fields))
// 	values := make([]interface{}, 0, len(fields)+1)

// 	// Prepare the SET clause with positional placeholders
// 	i := 1
// 	for field, value := range fields {
// 		setClauses = append(setClauses, field+" = $"+strconv.Itoa(i))
// 		values = append(values, value)
// 		i++
// 	}

// 	query := "UPDATE ticket_event_ticket SET " + strings.Join(setClauses, ", ") + " WHERE ticket_num = $" + strconv.Itoa(i)
// 	values = append(values, ticketNum)

// 	// Execute the update query with dynamic values
// 	_, err := pg.DB.ExecContext(ctx, query, values...)
// 	if err != nil {
// 		return false, errors.New("failed to update the ticket object: " + err.Error())
// 	}

// 	return true, nil
// }

// func (pg *LocalPgClient) CreateTicketObject(fields map[string]interface{}) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel() // Ensure context is canceled when the function returns

// 	if len(fields) == 0 {
// 		return errors.New("no fields provided for creating a payment object")
// 	}
// 	// Ensure the request_and_response field is set to an empty JSON object if not provided

// 	// Prepare the INSERT INTO query components
// 	columns := make([]string, 0, len(fields))
// 	placeholders := make([]string, 0, len(fields))
// 	values := make([]interface{}, 0, len(fields))

// 	// Build the columns, placeholders, and values slices
// 	i := 1
// 	for field, value := range fields {

// 		columns = append(columns, field)
// 		placeholders = append(placeholders, "$"+strconv.Itoa(i))
// 		values = append(values, value)
// 		// log.Debugf("Column: %s, Placeholder: $%d, Value: %v", field, i, value)

// 		i++
// 	}

// 	// Construct the final INSERT INTO query
// 	query := "INSERT INTO ticket_event_ticket (" + strings.Join(columns, ", ") + ") VALUES (" + strings.Join(placeholders, ", ") + ")"

// 	// Log the final query and parameters
// 	log.Debugf("Final SQL Query: %s", query)
// 	log.Debugf("Query Parameters: %v", values)

// 	// Execute the query
// 	res, err := pg.DB.ExecContext(ctx, query, values...)
// 	if err != nil {
// 		log.Error("Failed to create the ticket object:", err)
// 		return fmt.Errorf("failed to create the ticket object: %w", err)
// 	}

// 	// Log the result of the execution
// 	rowsAffected, _ := res.RowsAffected()
// 	log.Infof("Ticket object creation successful. Rows affected: %d", rowsAffected)

// 	return nil
// }


