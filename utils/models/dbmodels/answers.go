package dbmodels

type AnswerEntry struct {
	Index    int    `json:"index"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}



type Answer struct {
	UserID   int64         `json:"user_id"`
	SurveyID int64         `json:"survey_id"`
	Answers  []AnswerEntry `json:"answers"`  // Stored as JSONB in DB
}



