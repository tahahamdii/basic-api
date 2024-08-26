package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJson(r *http.Request, payload any ) error {
	if r.Body == nil {
		return fmt.Errorf("req body missing")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}