package ctrl_utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func CustomValidator(w http.ResponseWriter, reqBody []byte, target interface{}) error {
	// reqBody, _ := ioutil.ReadAll((*r).Body)
	fmt.Println("=> req Body : ", reqBody)
	if err := json.Unmarshal(reqBody, target); err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Printf("11")
		SendErrorResponse(w, http.StatusBadRequest, err.Error(), ValidationErrorType)
		return err
	}
	// if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
	// 	fmt.Println("11")
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return err
	// }

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(target)

	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error(), ValidationErrorType)
		fmt.Println("22")
		return err
	}
	// log.Warn().Str("validate opk", "validate ok").Msg("=>")
	return nil
}
