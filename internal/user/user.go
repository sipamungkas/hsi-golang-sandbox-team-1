package user

import (
	"time"

	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string         `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Password  string         `json:"password"`
	NIP       string         `gorm:"uniqueIndex; NOT NULL" json:"nip"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = cuid.New()
	return
}
