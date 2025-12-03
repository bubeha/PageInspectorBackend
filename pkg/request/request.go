package request

import (
	"encoding/json"
	"errors"
	"net/http"
)

func JSON(r *http.Request, i interface{}) error {
	if r.Body == nil {
		return errors.New("request body is empty")
	}

	r.Body = http.MaxBytesReader(nil, r.Body, 1_048_576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&i); err != nil {
		return errors.New("request body is not valid JSON")
	}

	if dec.More() {
		return errors.New("json: cannot unmarshal object into Go value of type httputil.Request")
	}

	return nil
}
