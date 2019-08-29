package helpers

import (
	"encoding/json"
	"net/http"
)

//ReturnSuccess returns succes response
func ReturnSuccess(w http.ResponseWriter, httpStatusCode int, data interface{}) error {

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.WriteHeader(httpStatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	return nil
}
