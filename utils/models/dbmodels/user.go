package dbmodels


type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`  // "admin" or "client"
	Email string `json:"email"`
}
