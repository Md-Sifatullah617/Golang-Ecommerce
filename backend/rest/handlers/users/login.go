package users

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var requestLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	usr := database.Find(requestLogin.Email, requestLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	cnf := config.GetConfig()
	accessToken, err := util.CreateJWT(cnf.JwtSecretKey, util.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, accessToken, http.StatusCreated)
}
