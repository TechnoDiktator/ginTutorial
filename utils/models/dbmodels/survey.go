package dbmodels


type Survey struct {
	ID          int64    `json:"id"`
	Heading     string   `json:"heading"`
	Description string   `json:"description"`
	Questions   []string `json:"questions"`  // Stored as array field in DB
}
