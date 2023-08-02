package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.com/searchEngineTumail/back/services"
	"github.com/go-chi/chi/v5"
)

var router = chi.NewRouter()

type ErrorResponse struct {
	Message string
}

func main() {
	router.Get("/search/match", searchMatch)
	http.ListenAndServe(":8080", router)
}

func searchMatch(w http.ResponseWriter, req *http.Request) {
	// se asigna a la cabecera el tipo json
	w.Header().Set("Content-Type", "application/json")
	// se obtienen los valores de los query params
	termPm := req.URL.Query().Get("term")
	startTimePm := req.URL.Query().Get("startTime")
	endTimePm := req.URL.Query().Get("endTime")
	maxResultsPm, err := strconv.Atoi(req.URL.Query().Get("maxResults"))
	if err != nil {
		w.WriteHeader(422)
		errorResponse := ErrorResponse{
			Message: err.Error(),
		}
		response, _ := json.Marshal(errorResponse)
		w.Write(response)
		return
	}

	fieldsPm := req.URL.Query().Get("fields")
	var fieldsList []string
	// si fields querypm viene vacío, se analizan todos los campos
	if fieldsPm == "" {
		fieldsList = []string{}
	} else {
		fieldsList = strings.Split(fieldsPm, "-")
	}

	// se llama a la función que tiene la lógica
	result, err := services.SearchMatch(termPm, startTimePm, endTimePm, maxResultsPm, fieldsList)
	if err != nil {
		w.WriteHeader(500)
		errorResponse := ErrorResponse{
			Message: err.Error(),
		}
		response, _ := json.Marshal(errorResponse)
		w.Write(response)
		return
	}

	// se convierte el map a json para poder retornar la respuesta en la API
	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(422)
		errorResponse := ErrorResponse{
			Message: err.Error(),
		}
		response, _ = json.Marshal(errorResponse)
		w.Write(response)
		return
	}

	// escribimos la respuesta json en el cuerpo de la respuesta HTTP para retornarla
	w.Write(response)
	return
}
