package entities

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Name string
}

func (e Entry) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"ID":        e.ID,
		"Name":      e.Name,
		"CreatedAt": e.CreatedAt,
	})
}
