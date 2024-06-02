package entity

type Order struct {
	BrothId   string `json:"brothId"`
	ProteinId string `json:"proteinId"`
}

func (*Order) TableName() string {
	return "orders"
}

func NewOrder(brothId, proteinId string) *Order {
	return &Order{
		BrothId:   brothId,
		ProteinId: proteinId,
	}
}
