package ctrl_utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func CustomBodyValidator(w http.ResponseWriter, r *http.Request, target interface{}) (error, uint) {
	// TODO : protect against very big bodys !
	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		return err, http.StatusBadRequest
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(target)

	if err != nil {
		return err, http.StatusBadRequest
	}

	return nil, 0
}
