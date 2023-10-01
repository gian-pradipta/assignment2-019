package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"rest_api_order/internal/controllers/validation/general"
	"rest_api_order/internal/repository/models/item"
	"rest_api_order/internal/repository/models/order"
)

type ItemValidator struct {
}

func (v *ItemValidator) isSingleJSONComplete(item item.Item) bool {

	var result bool = true
	result = result && !(item.Code == "")
	result = result && !(item.Description == "")
	result = result && !(item.Quantity == 0)

	return result
}

func (v *ItemValidator) DoDuplicateItemsExistInJSON(items []item.Item) error {
	var err error
	if items == nil {
		return err
	}
	var duplicationExistence bool = false
	var itemCodes []string = make([]string, len(items))
	for _, item := range items {
		if !general.IsInArray(item.Code, itemCodes) {
			itemCodes = append(itemCodes, item.Code)
		} else {
			duplicationExistence = true
			break
		}
	}
	if duplicationExistence {
		err = errors.New("Duplicate item codes detected")
	}
	return err
}

func (v *ItemValidator) DoesDuplicateExistInDB(items []item.Item) error {
	var err error
	if item.DoDuplicatesExist(items) {
		err = errors.New("code id is a duplicate")
	}
	return err
}

func (v *ItemValidator) IsJSONComplete(jsonByte []byte) (*[]item.Item, error) {
	var err error
	var validationStructs []item.Item

	err = json.Unmarshal(jsonByte, &validationStructs)
	if err != nil {
		return nil, err
	}
	var valid bool = true
	for _, validationStruct := range validationStructs {
		valid = valid && !(validationStruct.Code == "")
		valid = valid && !(validationStruct.Description == "")
		valid = valid && !(validationStruct.Quantity == 0)
	}

	if !valid {
		err = errors.New("Incomplete JSON for items")
		return nil, err
	}
	return &validationStructs, err
}

func (v *ItemValidator) isJSONCompletePartial(jsonByte []byte) (*order.Order, error) {
	var err error
	var validationMaps []map[string]interface{}
	json.Unmarshal(jsonByte, &validationMaps)

	var allowedParams []string = []string{"code", "description", "quantity"}
	var paramsAvailable []string = make([]string, len(allowedParams))
	for _, validationMap := range validationMaps {
		for key := range validationMap {
			if !general.IsInArray(key, allowedParams) {
				fmt.Println(key)
				fmt.Println(validationMap[key])
				err = errors.New("Invalid param on JSON")
				return nil, err
			}
			paramsAvailable = append(paramsAvailable, key)
		}
	}

	return nil, err
}

func (v *ItemValidator) IsPartialJSONValid(jsonByte []byte) error {
	var validationStructs []item.Item

	err := json.Unmarshal(jsonByte, &validationStructs)

	if err != nil {
		return err
	}
	for _, validationStruct := range validationStructs {
		if validationStruct.Code == "" {
			err = errors.New("Item JSON must include code")
			return err
		}
	}

	_, err = v.IsJSONComplete(jsonByte)
	if err == nil {
		return err
	}

	if _, err = v.isJSONCompletePartial(jsonByte); err != nil {
		return err
	}

	return err

}
