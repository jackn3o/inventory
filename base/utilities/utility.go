package utility

import (
	"encoding/json"
	"net/http"
)

type meta struct {
	Code     int         `json:"code,omitempty"`
	Status   string      `json:"status,omitempty"`
	Messages interface{} `json:"message,omitempty"`
}

type jsonResponseDTO struct {
	Data interface{} `json:"data,omitempty"`
	Meta meta        `json:"meta,omitempty"`
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

// WriteCreated as 201 common response
// For major POST with related info return
func (u *Utility) WriteCreated(value interface{}) {
	statusCode := http.StatusCreated

	u.WriteJSON(value, statusCode)
}

// WriteNoContent as 204 common response
// For major PUT and Patch without related info return
func (u *Utility) WriteNoContent(value interface{}) {
	statusCode := http.StatusNoContent

	u.WriteJSON(value, statusCode)
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

// WriteJSONMeta return JSON format data to client
func (u *Utility) WriteJSONMeta(value interface{}, statusCode int) {
	meta, err := json.Marshal(jsonResponseDTO{
		Meta: meta{
			Code:     statusCode,
			Status:   http.StatusText(statusCode),
			Messages: value,
		},
	})

	if err != nil {
		http.Error(u.Writer, err.Error(), http.StatusInternalServerError)
	} else {
		u.Writer.Write(meta)
	}
}

// NotImplemented as a utitlies
func (u *Utility) NotImplemented() {
	u.Writer.Write([]byte("Not Implemented"))
}
