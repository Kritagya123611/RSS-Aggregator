package auth

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type SignupRequest struct {
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var body SignupRequest
	// now trying to decode the json body which is recieved
	err:=json.NewDecoder(r.Body).Decode(&body);
	if err!=nil{
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}else{
		fmt.Printf("Email: %s, Password: %s\n", body.EMAIL, body.PASSWORD)
	}
}

