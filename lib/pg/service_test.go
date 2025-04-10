package pg

import (
	"testing"

	"github.com/TechnoDiktator/ginTutorial/serviceinit"
	"github.com/TechnoDiktator/ginTutorial/utils/models/dbmodels"
)

func Test_CreateSurvey(t *testing.T) {
	// Initialize all required services

	var s dbmodels.Survey

	s.Description = "We would love your feedback on our recent service."
	s.Heading = "Customer Satisfaction Survey"

	quesions := []string{
		"How satisfied were you with our service?",
		"Was our staff helpful and courteous?",
		"Would you recommend us to others?",
		"Any additional comments or suggestions?",
	}

	s.Questions = quesions

	localServiceInit, err := serviceinit.LocalServerInit()
	//initialize the localservice


	if err != nil {
		t.Fatalf("Connection to database Failed")
	}

	localServiceInit.PgService.




}
