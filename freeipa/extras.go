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

type Primitives interface {
	int | bool | string
}

type IPAPrim[T Primitives] struct {
	Value T
}

func (a *IPAPrim[T]) UnmarshalJSON(data []byte) error {
	var result T
	var resultStr string
	var multi []string
	err := json.Unmarshal(data, &result)
	if err != nil {
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
		var postConv any
		switch any(result).(type) {
		case bool:
			postConv, err = strconv.ParseBool(resultStr)
			if err != nil {
				return err
			}
		case int:
			postConv, err = strconv.ParseInt(resultStr, 10, 32)
			if err != nil {
				return err
			}
		}
		result = T(postConv)

	}

	*a = IPAPrim[T]{Value: result}
	return nil
}

type IPABool bool

func (a *IPABool) UnmarshalJSON(data []byte) error {
	var result bool
	var resultStr string
	var multi []string
	err := json.Unmarshal(data, &result)
	if err != nil {
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
