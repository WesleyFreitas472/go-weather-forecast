package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UnmarshalPayload(r *http.Response, model interface{}) error {
	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bodyBytes, &model)
}
