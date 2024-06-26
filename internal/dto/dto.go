package dto

type Broth struct {
	Id            string `json:"id"`
	ImageInactive string `json:"imageInactive"`
	ImageActive   string `json:"imageActive"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
}

type Protein struct {
	Id            string `json:"id"`
	ImageInactive string `json:"imageInactive"`
	ImageActive   string `json:"imageActive"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
}

type OrderRequest struct {
	BrothId   string `json:"brothId"`
	ProteinId string `json:"proteinId"`
}

type OrderResponse struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type OrderId struct {
	OrderId string `json:"orderId"`
}

type Error struct {
	Message string `json:"message"`
}
