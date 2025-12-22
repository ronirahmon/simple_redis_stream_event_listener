package helpers

import "encoding/json"

func Bind(data any, bind any) error {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}

	json.Unmarshal(jsonString, bind)
	return nil

}
