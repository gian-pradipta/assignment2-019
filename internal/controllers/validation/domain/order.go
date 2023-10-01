package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"rest_api_order/internal/controllers/validation/general"
	"rest_api_order/internal/repository/models/order"
)

type OrderValidator struct {
}

func (v *OrderValidator) isJSONComplete(jsonByte []byte) error {
	var requiredLen int = 3
	var validationMap map[string]interface{}
	var err error
	err = json.Unmarshal(jsonByte, &validationMap)
	if len(validationMap) != requiredLen {
		err = errors.New("incomplete JSON: JSON requires customer_name, items, and ordered_at")
	}
	return err
}

func (v *OrderValidator) ContentValidation(jsonByte []byte) (*order.Order, error) {
	var validationStruct order.Order
	var validationMap map[string]interface{}
	var err error
	err = json.Unmarshal(jsonByte, &validationMap)
	for key := range validationMap {
		if !general.IsInArray(key, []string{"customer_name", "ordered_at", "items"}) {
			err = errors.New("Invalid JSON")
			return &validationStruct, err
		}
	}
	err = json.Unmarshal(jsonByte, &validationStruct)
	return &validationStruct, err
}

func (v *OrderValidator) IsJSONComplete(jsonByte []byte) (*order.Order, error) {
	var validationStruct order.Order
	err := json.Unmarshal(jsonByte, &validationStruct)
	if err != nil {
		return nil, err
	}

	var result bool = true

	result = result && !(validationStruct.CustomerName == "")
	result = result && !(validationStruct.OrderedAt == "")
	result = result && !(validationStruct.Items == nil)
	if !result {
		err = errors.New("Incomplete JSON detected")
		return nil, err
	}
	return &validationStruct, err

}

func (v *OrderValidator) IsJSONCompletePartial(jsonByte []byte) (*order.Order, error) {
	var err error
	var validationMap map[string]interface{}
	var validationStruct order.Order
	json.Unmarshal(jsonByte, &validationMap)
	json.Unmarshal(jsonByte, &validationStruct)

	var allowedParams []string = []string{"customer_name", "ordered_at", "items"}
	var paramsAvailable []string = make([]string, len(allowedParams))

	for key := range validationMap {
		if !general.IsInArray(key, allowedParams) {
			fmt.Println(key)
			err = errors.New("Invalid param on JSON1")
			return &validationStruct, err
		}
		paramsAvailable = append(paramsAvailable, key)
	}

	return &validationStruct, err
}
