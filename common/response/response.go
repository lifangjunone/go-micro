package response

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Code    *int
	Message string
	Data    interface{}
}

func Success(w http.ResponseWriter, data interface{}) {
	c := 10000
	resp := Data{
		Code:    &c,
		Message: "Success",
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}

func Failed(w http.ResponseWriter, data interface{}) {
	c := 10001
	resp := Data{
		Code:    &c,
		Message: "Failed",
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}

func ParamMiss(w http.ResponseWriter, data interface{}) {
	c := 10002
	resp := Data{
		Code:    &c,
		Message: "Param was missing",
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	respBytes, _ := json.Marshal(resp)
	w.Write(respBytes)
}
