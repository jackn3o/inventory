package utility

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Code    int         `json:"code,omitempty"`
	Status  string      `json:"status,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

type jsonResponseDTO struct {
	Data  interface{}   `json:"data,omitempty"`
	Error errorResponse `json:"meta,omitempty"`
}

// Utility is common helper
type Utility struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

// New is a constructor for creating utility helper
func New(writer http.ResponseWriter, req *http.Request) *Utility {
	return &Utility{
		Writer:  writer,
		Request: req,
	}
}

// WriteJSON return JSON format data to client, default 200
func (u *Utility) WriteJSON(value interface{}, vars ...int) error {
	statusCode := http.StatusOK
	if len(vars) > 0 {
		statusCode = vars[0]
	}
	data, err := json.Marshal(jsonResponseDTO{
		Data: value,
	})

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
	errResponse, err := json.Marshal(jsonResponseDTO{
		Error: errorResponse{
			Code:    statusCode,
			Status:  http.StatusText(statusCode),
			Message: value,
		},
	})

	if err != nil {
		http.Error(u.Writer, err.Error(), http.StatusInternalServerError)
	} else {
		u.Writer.Write(errResponse)
	}
}

// NotImplemented as a utitlies
func (u *Utility) NotImplemented() {
	u.Writer.Write([]byte("Not Implemented"))
}
