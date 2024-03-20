package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseBody reads the request body and parses it into the provided interface
func ParseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &x)
	if err != nil {
		return err
	}

	return nil
}