package repository

import (
	"hoc-gin/internal/models"
	"log"

	"gorm.io/gorm"
)

type SQLUserRepository struct {
	db *gorm.DB
}

func NewSQLUserRepository(DB *gorm.DB) UserRepository {
	return &SQLUserRepository{
		db: DB,
	}
}

func (ur *SQLUserRepository) Create(user *models.User) {
	log.Println("Create")
}

func (ur *SQLUserRepository) FindById(user *models.User, id int) error {
	if err := ur.db.First(user, id).Error; err != nil {
		return err
	}

	return nil
}
