package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(req *http.Request, res interface{}) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(res)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, res interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(res)
	PanicIfError(err)
}