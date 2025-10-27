package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Category struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Slug      string         `gorm:"unique;not null" json:"slug"`
	Name      string         `gorm:"not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Tag struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Slug      string         `gorm:"unique;not null" json:"slug"`
	Name      string         `gorm:"not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Component struct {
	ID              uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Slug            string         `gorm:"unique;not null" json:"slug"`
	Name            string         `gorm:"not null" json:"name"`
	Description     string         `json:"description"`
	CategoryID      uuid.UUID      `gorm:"not null" json:"-"`
	Category        Category       `gorm:"foreignKey:CategoryID" json:"category"`
	CodeJSX         string         `gorm:"type:text;not null" json:"code_jsx"`
	CodeCSS         string         `gorm:"type:text" json:"code_css,omitempty"`
	PropsDefinition datatypes.JSON `json:"props_definition"`
	UserID          uuid.UUID      `gorm:"not null" json:"user_id"`
	Tags            []*Tag         `gorm:"many2many:component_tags;" json:"tags,omitempty"`
	Status          string         `json:"status"`
	ApprovalStatus  string         `json:"approval_status"`
	ReviewerID      uuid.UUID      `json:"reviewer_id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
