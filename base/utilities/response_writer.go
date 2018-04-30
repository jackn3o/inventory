package utility

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Code     int         `json:"code,omitempty"`
	Status   string      `json:"status,omitempty"`
	Messages interface{} `json:"messages,omitempty"`
}

type jsonErrorResponseDTO struct {
	Error *errorResponse `json:"meta,omitempty"`
}

// WriteJSON return JSON format data to client, default 200
func (u *Utility) WriteJSON(value interface{}, vars ...int) error {
	statusCode := http.StatusOK
	if len(vars) > 0 {
		statusCode = vars[0]
	}
	data, err := json.Marshal(value)

	if err != nil {
		http.Error(u.Writer, err.Error(), http.StatusInternalServerError)
	} else {
		u.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if statusCode != http.StatusOK {
			u.Writer.WriteHeader(statusCode)
		}

		u.Writer.Write(data)
	}

	return err
}

// WriteJSONError return JSON error to client
func (u *Utility) WriteJSONError(value interface{}, statusCode int) {
	errResponse, err := json.Marshal(jsonErrorResponseDTO{
		Error: &errorResponse{
			Code:     statusCode,
			Status:   http.StatusText(statusCode),
			Messages: value,
		},
	})

	if err != nil {
		http.Error(u.Writer, err.Error(), http.StatusInternalServerError)
	} else {
		u.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		u.Writer.WriteHeader(statusCode)
		u.Writer.Write(errResponse)
	}
}

// NotImplemented as a utitlies
func (u *Utility) NotImplemented() {
	u.Writer.Write([]byte("Not Implemented"))
}
