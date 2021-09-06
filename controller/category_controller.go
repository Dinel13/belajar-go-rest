package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	// Create params karena use lib httprouter
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
