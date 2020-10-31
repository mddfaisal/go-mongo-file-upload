package uploadcontoller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// upload upload json
type upload struct {
	Muid     string `json:"muid"`
	Filename string `json:"filename"`
	Mimetype string `json:"mimetype"`
	Data     string `json:"data"`
}

// Upload upload files
func Upload(w http.ResponseWriter, r *http.Request) {
	var data upload
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
