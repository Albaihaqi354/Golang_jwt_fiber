package helpers

import (
	"encoding/json"
	"net/http"
)

type ResponWithData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponWithOutData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, code int, message string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var response any
	status := "Success"

	if code >= 400 {
		status = "Failed"
	}

	if payload != nil {
		response = ResponWithData{
			Status:  status,
			Message: message,
			Data:    payload,
		}
	} else {
		response = ResponWithOutData{
			Status:  status,
			Message: message,
		}
	}

	res, _ := json.Marshal(response)
	w.Write(res)
}
