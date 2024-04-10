package ctrl_utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func CustomValidator(w http.ResponseWriter, r *http.Request, target *interface{}) error {
	reqBody, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(reqBody, target); err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		SendErrorResponse(w, http.StatusBadRequest, err.Error(), ValidationErrorType)
		return err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(target)

	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, err.Error(), ValidationErrorType)
		return err
	}
	// log.Warn().Str("validate opk", "validate ok").Msg("=>")
	return nil
}
