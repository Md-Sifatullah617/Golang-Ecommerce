package products

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pID, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give me a valid product id", http.StatusBadRequest)
		return
	}

	product, err := h.svc.Get(pID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if product == nil {
		util.SendData(w, "No data found with that product id", http.StatusNotFound)
		return
	}

	util.SendData(w, product, http.StatusOK)
}
