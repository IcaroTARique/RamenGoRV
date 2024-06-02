package entity

type Answer struct {
	Id          string `json:"id"`
	BrothId     string `json:"broth_id"`
	ProteinId   string `json:"protein_id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (*Answer) TableName() string {
	return "answer"
}
