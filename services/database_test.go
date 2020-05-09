package services

import (
	"testing"

	"github.com/jinzhu/gorm"
)

// TestModel only for test
type TestModel struct {
	gorm.Model
	Code  string
	Price uint
}

func TestDB(t *testing.T) {
	DB.AutoMigrate(&TestModel{})
	DB.Create(&TestModel{Code: "L1212", Price: 1000})
	var t1 TestModel
	DB.First(&t1, 1)
	// DB.Delete(&t1)
}
