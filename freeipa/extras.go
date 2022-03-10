package freeipa

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type IPATime time.Time

func (a *IPATime) UnmarshalJSON(data []byte) error {
	var result time.Time
	var finalStr string
	var single map[string]string
	var multi []map[string]string
	err := json.Unmarshal(data, &single)
	if err != nil {
		err = json.Unmarshal(data, &multi)
		if err != nil {
			return err
		}
		if len(multi) == 0 {
			return fmt.Errorf("DNSName mapping was empty")
		}
		single = multi[0]
	}
	// Extract just one value
	for _, v := range single {
		finalStr = v
		break
	}
	// 2024 03 01 21 34 41 Z
	//"Mon Jan _2 15:04:05 2006"
	result, err = time.Parse("20060102150405Z", finalStr)
	if err != nil {
		return fmt.Errorf("Unable to parse time %s: %w", finalStr, err)
	}
	*a = IPATime(result)
	return nil
}

type DNSName string

func (a *DNSName) UnmarshalJSON(data []byte) error {
	var result string
	var single map[string]string
	var multi []map[string]string
	err := json.Unmarshal(data, &single)
	if err != nil {
		err = json.Unmarshal(data, &multi)
		if err != nil {
			return err
		}
		if len(multi) == 0 {
			return fmt.Errorf("DNSName mapping was empty")
		}
		single = multi[0]
	}
	// Extract just one value
	for _, v := range single {
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
			postConv, err = strconv.Atoi(resultStr)
			if err != nil {
				return err
			}
		default:
			postConv = resultStr
		}
		result = postConv.(T)
	}

	*a = IPAPrim[T]{Value: result}
	return nil
}
