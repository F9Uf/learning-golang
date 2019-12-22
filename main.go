package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	// :8080 => 0.0.0.0:8080
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}

type addRequest struct {
	A int `json:"x"`
	B int `json:"y"`
}

type addResponse struct {
	Result int `json:"res,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// pare request
	var reqBody addRequest
	json.NewDecoder(r.Body).Decode(&reqBody)

	// calculate
	result := reqBody.A + reqBody.B

	// send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addResponse{
		Result: result,
	})
}
