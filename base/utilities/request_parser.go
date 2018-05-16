package utility

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
)

func (u *Utility) parseJSON(value interface{}) (int, error) {
	body, err := ioutil.ReadAll(u.Request.Body)
	if err != nil {
		return http.StatusBadRequest, err
	}
	if err := json.Unmarshal(body, &value); err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

// Unmarshal is function to parse json from request header
func (u *Utility) Unmarshal(value interface{}) (int, error) {
	canonicalName := http.CanonicalHeaderKey("content-type")
	headerValue := u.Request.Header.Get(canonicalName)
	if headerValue != "" {
		if strings.Contains(headerValue, "application/json") {
			return u.parseJSON(value)
		}
	}
	return http.StatusBadRequest, nil
}

// UnmarshalWithValidation is function to validate parse json from request header
// Must have a validate function for the provide receiver
func (u *Utility) UnmarshalWithValidation(value interface{}) interface{} {
	if _, err := u.Unmarshal(value); err != nil {
		return err
	}
	govalidator.SetFieldsRequiredByDefault(true)
	_, err := govalidator.ValidateStruct(value)

	if err != nil {
		errs := err.(govalidator.Errors).Errors()
		errWriter := ErrorWriter{}
		for _, e := range errs {
			v := e.(govalidator.Error)

			if v.Validator == "required" && !v.CustomErrorMessageExists {
				v.Err = errors.New(v.Name + " is required")
			}

			errWriter.Add(v.Name, v.Err.Error())
		}

		// make result in validations object
		result := make(map[string]interface{})
		result["validations"] = errWriter.Errors()

		return result
	}

	return nil
}
