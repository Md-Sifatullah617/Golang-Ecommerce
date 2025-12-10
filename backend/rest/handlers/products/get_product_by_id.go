package products

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pID, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	product := database.Get(pID)
	if product == nil {
		util.SendData(w, "No data found with that product id", 404)
		return
	}

	util.SendData(w, product, 200)
}
