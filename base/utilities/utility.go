package utility

import "net/http"

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
