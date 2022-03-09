package freeipa

import (
	"encoding/json"
	"strconv"
)

type DNSName string

func (a *DNSName) UnmarshalJSON(data []byte) error {
	var result string
	var initialResult map[string]string
	if err := json.Unmarshal(data, &initialResult); err != nil {
		return err
	}
	// Extract just one value
	for _, v := range initialResult {
		result = v
		break
	}
	*a = DNSName(result)
	return nil
}

type IPAPrim<T> T

type IPABool bool

func (a *IPABool) UnmarshalJSON(data []byte) error {
	var result bool
	var resultBool bool
	var resultStr string
	var multi []string
	err := json.Unmarshal(data, &resultBool)
	if err == nil {
		result = resultBool
	} else {
		err = json.Unmarshal(data, &multi)
		if err != nil {
			err = json.Unmarshal(data, &resultStr)
			if err != nil {
				return err
			}
		} else {
			// Extract just one value
			for _, v := range multi {
				resultStr = v
				break
			}
		}

		result, err = strconv.ParseBool(resultStr)
		if err != nil {
			return err
		}
	}

	*a = IPABool(result)
	return nil
}
