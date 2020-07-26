package mongocontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"services/email"
	"services/utils"
)

// GetEmail get email
func GetEmail(w http.ResponseWriter, r *http.Request) {
	var e email.Email
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println("Error in json decode.")
		fmt.Println(utils.Trace())
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
	jsonData := ""
	responseCode := http.StatusOK
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		panic(err)
	}
	if e.Muid == "" || e.Subject == "" || e.HTML == "" {
		responseCode = http.StatusBadRequest
		jsonData = `{"_id": "", "msg": "No Content"}`
	} else {
		e.DecodeHTML()
		e.DecodeSubject()
		if e.Subject == "" || e.HTML == "" {
			responseCode = http.StatusBadRequest
			jsonData = `{"_id": "", "msg": "No Content"}`
		} else {
			e.Create()
			jsonData = `{"_id": "` + e.ID + `", "msg": "OK"}`
			responseCode = http.StatusCreated
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	w.Write([]byte(jsonData))
}

// UpdateEmail update email
func UpdateEmail(w http.ResponseWriter, r *http.Request) {}

// DeleteEmail delete email
func DeleteEmail(w http.ResponseWriter, r *http.Request) {}
