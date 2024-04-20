package models

type Admins struct {
	ID                int    `json:"id" gorm:"primary_key"`
	Email             string `json:"email" gorm:"unique;not null"`
	Password          string `json:"password" gorm:"not null"`
	FirstName         string `json:"first_name" gorm:"not null"`
	LastName          string `json:"last_name" gorm:"not null"`
	Phone             string `json:"phone" gorm:"not null"`
	IsEmailVerified   bool   `json:"is_email_verified" gorm:"default:false"`
	IsPhoneVerified   bool   `json:"is_phone_verified" gorm:"default:false"`
	IsPasswordChanged bool   `json:"is_password_changed" gorm:"default:false"`
	IsPhoneChanged    bool   `json:"is_phone_changed" gorm:"default:false"`
	IsEmailChanged    bool   `json:"is_email_changed" gorm:"default:false"`
	IsDeleted         bool   `json:"is_deleted" gorm:"default:false"`
	CreatedAt         int    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         int    `json:"updated_at" gorm:"autoUpdateTime"`
}