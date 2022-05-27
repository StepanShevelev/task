package db

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string     `json:"name" db:"name"`
	Category []Category `json:"category" db:"category" gorm:"many2many:user_categories;"`
	Pet      []Pet      `json:"pet" db:"pet"`
}

type Category struct {
	gorm.Model
	Name string `json:"name" db:"name"`
}

type Pet struct {
	gorm.Model
	Name    string `json:"name" db:"name"`
	Counter int    `json:"counter" db:"counter"`
	UserID  int    `json:"user_id" db:"user_id"`
}

func (p *Pet) ShowSkill() {
	fmt.Println("питомец показывает что умеет")
	sCounter := p.Counter
	sCounter += 1
	Client.First(&p)
	p.Counter = sCounter
	Client.Save(&p)
}
