package utility

import (
	"encoding/json"
	"net/http"
)

type meta struct {
	Code     int    `json:"code"`
	Status   string `json:"status"`
	Messages string `json:"message,omitempty"`
}

type jsonResponseDTO struct {
	Data interface{} `json:"data,omitempty"`
	Meta meta        `json:"meta,omitempty"`
	Info string      `json:"info,omitempty"`
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

// WriteStatus return code and status to client
func (u *Utility) WriteStatus(code int, vars ...string) {
	status := http.StatusText(code)
	if len(vars) > 0 {
		status = vars[0]
	}

	u.Writer.WriteHeader(code)
	u.Writer.Write([]byte(status))
}

// WriteJSON return JSON format data to client
func (u *Utility) WriteJSON(value interface{}, vars ...int) error {
	statusCode := http.StatusOK
	if len(vars) > 0 {
		statusCode = vars[0]
	}
	data, err := json.Marshal(jsonResponseDTO{
		Data: value,
		Meta: meta{
			Code:   statusCode,
			Status: http.StatusText(statusCode),
		},
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
