package domain

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
	Birthday  time.Time
}
