package pg

import (
	// "dynamic-qr/internal/util/models"

	"context"
	"time"

	"github.com/TechnoDiktator/ginTutorial/utils/models/dbmodels"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

//

type LocalPgClient struct {
	DB *sqlx.DB
}

func NewLocalPGService(db *sqlx.DB) LocalPgService {
	logrus.Info("NewPgService initialized")
	return &LocalPgClient{DB: db}
}

type LocalPgService interface {
	GetDB() *sqlx.DB // GETTER AND SETTER FUCNTION FOR DIFFERENT TABLES

	CreateSurvey(survey dbmodels.Survey) error
}

// GetDB returns the database connection
func (pg *LocalPgClient) GetDB() *sqlx.DB {
	return pg.DB
}

func (pg *LocalPgClient) CreateSurvey(survey dbmodels.Survey) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := `
		INSERT INTO surveys (heading, description, questions)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var insertedID int
	err := pg.DB.QueryRowxContext(ctx, query, survey.Heading, survey.Description, pq.Array(survey.Questions)).Scan(&insertedID)
	if err != nil {
		logrus.Errorf("Failed to insert survey: %v", err)
		return err
	}

	logrus.Infof("Survey inserted with ID: %d", insertedID)
	return nil
}

// type LocalPgService interface {
// 		// GETTER AND SETTER FUCNTION FOR DIFFERENT TABLES
// 		// TICKET
// 		GetCashTicketObjectByTicketNum(string) (*localdqrdbmodels.TicketEventTicket, error)
// 		GetTicketObjectByTicketNum(string) (*localdqrdbmodels.TicketEventTicket, error)        // Ticket Table
// 		GetTicketObjectByDqrTransactionId(string) (*localdqrdbmodels.TicketEventTicket, error) // Ticket Table
// 		UpdateTicketObject(string, map[string]interface{}) (bool, error)                       // Ticket Table
// 		GetLastTicketObject(int) (*localdqrdbmodels.TicketEventTicket, error)
// 		GetLastTicketClosedSameAsCurrent() (string, error)
// 		CreateTicketObject(map[string]interface{}) error
// 		// PAYMENT
// 		CreatePaymentObjectAndReturnIdDirectly(float64, string, string, string, string) (int64, error)
// 		GetPaymentObjectFromTicketNum(int) (*localdqrdbmodels.Payment, error)           // New Payment Table
// 		GetPaymentObjectFromDqrTransactionId(string) (*localdqrdbmodels.Payment, error) // New Payment Table
// 		UpdatePaymentObject(string, map[string]interface{}) error                       // New Payment Table
// 		UpdatePaymentObjectStatusDirectly(string, string) error
// 		CreatePaymentObject(map[string]interface{}) error // New Payment Table
// 		FetchPaymentObjUsingTicketId(int) (*localdqrdbmodels.Payment, error)
// 		FetchPaymentObjUsingId(int64) (*localdqrdbmodels.Payment, error)
// 		FetchPaymentObjUsingTicketIdAndAgent(int, string) (*localdqrdbmodels.Payment, error)
// 		FetchPaymentObjInPendingStateTicketIdAndDeviceType(int, string) ([]localdqrdbmodels.Payment, error)
// 		CheckPaymentAlreadyDoneThroughNotAllowedAgents(int) ([]localdqrdbmodels.Payment, error)
// 		UpdatePaymentObjectStatus(string, bool) (bool, error)
// 		GetToken(key string) (bool, error)
// 		CreatePaymentObjectId() (string, error)
// 		AllLocalDqrRequest(map[string]interface{}) (int64, error) // New Payment Table
// 		CreatePaymentObjectAndGetId(map[string]interface{}) (int64, error)
// 		UpdatePaymentObjectStatusToPassedArg(string, bool, string) (bool, error)
// 		UpdatePaymentObjectStatusToCancel(string, bool) (bool, error)
// 		FetchPaymentObjUsingTicketIdAndStatus(int, string) ([]localdqrdbmodels.Payment, error)
// 		FetchAllPaymentObjectsUsingState(string) ([]localdqrdbmodels.Payment, error)
// 		// GATE
// 		GetDefaultGate() (*localdqrdbmodels.Gate, error) // Gate Table
// 		GetAllGates() ([]localdqrdbmodels.Gate, error)
// 		// PARKZAP REQUEST AND RESPONSE
// 		CreateAllParkzapRequestObject(map[string]interface{}) (int64, error)
// 		GetAllParkzapRequestObject(string) (*localdqrdbmodels.AllParkzapRequestAndResponse, error)
// 		UpdateAllParkzapRequestUsingId(int64, map[string]interface{}) error
// 	}
