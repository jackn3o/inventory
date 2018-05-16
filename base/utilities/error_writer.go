package utility

import "fmt"

// ErrorsInfo is map set errors
type ErrorsInfo map[string]interface{}

func (e ErrorsInfo) Error() string {
	s := ""
	for k, v := range e {
		s += fmt.Sprintf("%s: %v\n", k, v)
	}
	return s
}

// ErrorWriter provides a data type to map error messages
type ErrorWriter struct {
	errors ErrorsInfo
}

// Add error message identified by key
func (e *ErrorWriter) Add(key string, value interface{}) {
	if e.errors == nil {
		e.errors = make(map[string]interface{})
	}
	if _, exists := e.errors[key]; !exists {
		e.errors[key] = value
	}
}

// Errors method return errors information stored in a map
func (e *ErrorWriter) Errors() interface{} {
	if e.errors == nil {
		return nil
	}
	// copy the internal errors' map
	result := make(map[string]interface{})
	for key, value := range e.errors {
		result[key] = value
	}
	return result
}
