package item

import (
	"errors"
	"rest_api_order/internal/repository/database"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Code        string `gorm:"unique" json:"code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var db *gorm.DB = database.New()

func DoesDuplicateExist(item Item) bool {
	var tempItem Item
	db.Find(&tempItem, "code = ?", item.Code)
	if tempItem.ID == 0 {
		return false
	}
	return true
}

func DoDuplicatesExist(items []Item) bool {
	var result bool = false
	for _, item := range items {
		result = result || DoesDuplicateExist(item)
	}
	return result
}

func UpdateItemOnCode(newItem *Item) error {
	var updatedItem Item
	db.First(&updatedItem, "code = ?", newItem.Code)
	if updatedItem.ID == 0 {
		return errors.New("No item with that code")
	}
	updatedItem.Quantity = newItem.Quantity
	updatedItem.Description = newItem.Description
	db.Save(&updatedItem)
	return nil
}

func InsertData(id uint, newItem *Item) (uint, error) {
	var err error
	newItem.OrderID = id
	if newItem.Description == "" || newItem.Quantity == 0 {
		err = errors.New("Incomplete Data Item")
		return newItem.ID, err
	}
	db.Create(newItem)
	return newItem.ID, err
}

func CompleteItem(incompleteItem *Item, completeItem *Item) {

	if incompleteItem.Description == "" {
		incompleteItem.Description = completeItem.Description
	}
	if incompleteItem.Quantity == 0 {
		incompleteItem.Quantity = completeItem.Quantity
	}
}

func UpdateItemOnCodePartial(newItem *Item) error {
	var updatedItem Item
	db.First(&updatedItem, "code = ?", newItem.Code)
	if updatedItem.ID == 0 {
		return errors.New("No item with that code")
	}
	CompleteItem(newItem, &updatedItem)
	updatedItem.Quantity = newItem.Quantity
	updatedItem.Description = newItem.Description
	db.Save(&updatedItem)
	return nil
}
