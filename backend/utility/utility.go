package utility

import (
	"encoding/json"
	"net/http"
)

func Interface2Object(origin interface{}, target interface{}) (err error) {
	byteData, err := json.Marshal(origin)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteData, &target)
}

func SetupResponseCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}
