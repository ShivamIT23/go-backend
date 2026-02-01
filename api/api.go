package api

import (
	"net/http"
	"encoding/json"
)

type CoinBalanceParams struct{
	Username string
}

type CoinBalanceResponse struct{
	// Code to send on response, usually 200 for success
	Code int8

	Balance int64
}

type Error struct{
	//same code for error, usually 500 for internal server error
	Code int8

	Message string
}

func writeError( w http.ResponseWriter , message string , code int8){

	resp := Error{
		Code : code,
		Message : message
	}

	w.Header().set("Content-type","application/json")

	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	requestErrorHandler = func(w http.ResponseWriter , err error){
		writeError(w,err.Error(),http.StatusBadRequest)
	}
	internalErrorHandler = func(w http.ResponseWriter){
		writeError(w,"An unExpected Error Occured",http.StatusInternalServerError)
	}
)
