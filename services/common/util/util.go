package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseJSON(r *http.Request, v interface{}) error {
	if r.Body == nil {
		return io.EOF
	}
	defer r.Body.Close() 

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}
	return nil
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
}

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf(`{"data": "%v"}`, data)))
}
