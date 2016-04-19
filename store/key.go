package store

import (
	"github.com/jinzhu/gorm"
	"github.com/solderapp/solder-api/model"
	"golang.org/x/net/context"
)

// GetKeys retrieves all available keys from the database.
func GetKeys(c context.Context) (*model.Keys, error) {
	return FromContext(c).GetKeys()
}

// CreateKey creates a new key.
func CreateKey(c context.Context, record *model.Key) error {
	return FromContext(c).CreateKey(record)
}

// UpdateKey updates a key.
func UpdateKey(c context.Context, record *model.Key) error {
	return FromContext(c).UpdateKey(record)
}

// DeleteKey deletes a key.
func DeleteKey(c context.Context, record *model.Key) error {
	return FromContext(c).DeleteKey(record)
}

// GetKey retrieves a specific key from the database.
func GetKey(c context.Context, id string) (*model.Key, *gorm.DB) {
	return FromContext(c).GetKey(id)
}