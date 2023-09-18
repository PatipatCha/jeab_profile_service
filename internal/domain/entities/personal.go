package entities

type Personal struct {
	ID        uint
	FirstName string `json:"firstname"`
	SurName   string `json:"surname"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
}
