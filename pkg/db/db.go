package db

import (
	"fmt"
	cfg "github.com/StepanShevelev/task/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *client

type client struct {
	*gorm.DB
}

func New(config *cfg.Config) (*client, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=Europe/Moscow",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("cant connect to db: %w", err)
	}

	tmp := &client{db}
	Client = tmp
	return tmp, nil
}

func (db *client) DbGetUserById(ID int) (*User, error) {
	user := NewUser()
	result := db.Where("id = ?", ID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (db *client) DbGetUsers() ([]User, error) {
	var users []User

	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (db *client) DbGetCategoryByID(ID int) (*Category, error) {
	cat := NewCategory()
	result := db.Where("id = ?", ID).First(&cat)
	if result.Error != nil {
		return nil, result.Error
	}
	return cat, nil
}

func (db *client) DbGetCategories() ([]Category, error) {
	var categories []Category

	result := db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (db *client) DbGetPetByID(ID int) (*Pet, error) {
	pet := NewPet()
	result := db.Where("id = ?", ID).First(&pet)
	if result.Error != nil {
		return nil, result.Error
	}
	return pet, nil
}

func (db *client) DbGetUserPetByID(id int) (*Pet, error) {

	pet := NewPet()
	result := db.First(&pet, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return pet, nil
}

func (db *client) DbGetPets() ([]Pet, error) {
	var pets []Pet

	result := db.Find(&pets)
	if result.Error != nil {
		return nil, result.Error
	}
	return pets, nil
}

func NewUser() *User {
	return &User{}
}

func NewPet() *Pet {
	return &Pet{}
}

func NewCategory() *Category {
	return &Category{}
}
func (db *client) SetDB() error {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Pet{})
	return nil
}
