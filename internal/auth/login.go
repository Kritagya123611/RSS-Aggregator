package auth

import(
	"encoding/json"
	"net/http"
	"fmt"
)

type LoginRequest struct{
	NAME string `json:"name"`
	EMAIL string `json:"email"`
	PASSWORD string `json:"password"`
}
 
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var body LoginRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println("Login endpoint hit")
	fmt.Printf("Name: %s, Email: %s, Password: %s\n", body.NAME, body.EMAIL, body.PASSWORD)
	fmt.Println("User logged in successfully")

	w.Write([]byte("Login successful!"))
}