package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string         `gorm:"primary_key"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (baseModel *BaseModel) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewString()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
