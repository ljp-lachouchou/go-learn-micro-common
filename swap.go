package common

import "encoding/json"

func SwapTo(req, product interface{}) (err error) {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bytes, product); err != nil {
		return err
	}
	return nil
}
