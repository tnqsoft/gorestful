package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Contact struct {
	ID    int    `json:"id" gorm:"primary_key:yes"`
	Name  string `json:"name" gorm:"type:varchar(255);not null"`
	Email string `json:"email" gorm:"type:varchar(100)"`
	Phone string `json:"phone" gorm:"type:varchar(30)"`
}

// func (e *Contact) Disable() {
// 	e.Status = false
// }

// func (p *Employee) Enable() {
// 	p.Status = true
// }

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Contact{})
	return db
}
