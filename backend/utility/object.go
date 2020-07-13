package utility

import (
    "encoding/json"
)
    
func Interface2Object(origin interface{}, target interface{}) (err error) {
	byteData, err := json.Marshal(origin)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteData, &target)
}