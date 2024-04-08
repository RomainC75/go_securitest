package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/api/dto/requests"
)

func HandleSignupUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req requests.SignupRequest
	json.Unmarshal(reqBody, &req)
	// json.NewEncoder(w).Encode(post)

	fmt.Println("=> <", req)
}
