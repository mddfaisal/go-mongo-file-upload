package mongocontroller

import (
	"encoding/json"
	"fmt"
	"mongoservice/email"
	"net/http"
)

// GetEmail get email
func GetEmail(w http.ResponseWriter, r *http.Request) {
	var e email.Email
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println("Error in json decode.")
	}
	em, err := e.FindOne()
	w.Header().Set("Content-Type", "application/json")
	responseCode := map[bool]int{true: http.StatusOK, false: http.StatusNotFound}[err == nil]
	// response := map[bool][]byte{true: em, false: []byte(err.Error())}[err == nil]
	w.WriteHeader(responseCode)
	w.Write(em)
}

// NewEmail create new email
func NewEmail(w http.ResponseWriter, r *http.Request) {
	var e email.Email
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println("Error in json decode.")
	}
	id, err := e.New()
	if err != nil {
		fmt.Println(err)
	}
	json := `{"lastid": "` + id + `"}`
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(json))
}

// UpdateEmail update email
func UpdateEmail(w http.ResponseWriter, r *http.Request) {}

// DeleteEmail delete email
func DeleteEmail(w http.ResponseWriter, r *http.Request) {}
