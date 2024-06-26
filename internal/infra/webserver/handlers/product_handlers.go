package handlers

import (
	"encoding/json"
	"github.com/IcaroTARique/RamenGoRV/internal/dto"
	"github.com/IcaroTARique/RamenGoRV/internal/infra/port"
	"net/http"
)

type ProductHandler struct {
	ProductPersistence port.ProductInterface
}

func NewProductHandler(productPersistence port.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductPersistence: productPersistence}
}

func (p *ProductHandler) ListBroths(w http.ResponseWriter, r *http.Request) {
	broths, err := p.ProductPersistence.GetBroths()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.Error{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(broths); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.Error{Message: "Internal error"})
	}
	w.WriteHeader(http.StatusOK)
}

func (p *ProductHandler) ListProteins(w http.ResponseWriter, r *http.Request) {
	proteins, err := p.ProductPersistence.GetProtein()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.Error{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(proteins); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.Error{Message: "Internal error"})
	}
	w.WriteHeader(http.StatusOK)
}

func (p *ProductHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var OrderRequest dto.OrderRequest
	err := json.NewDecoder(r.Body).Decode(&OrderRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.Error{Message: err.Error()})
		return
	}
	res, err := p.ProductPersistence.CreateOrder(OrderRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.Error{Message: err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.Error{Message: "Internal error"})
	}
	w.WriteHeader(http.StatusCreated)
}
