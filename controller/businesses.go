package controller

import (
	"business/model"
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func GetBusinesses(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//получаем список всех пользователей
	businesses, err := model.GetAllBusinesses()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//возвращаем список клиенту в формате JSON
	err = json.NewEncoder(rw).Encode(businesses)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}