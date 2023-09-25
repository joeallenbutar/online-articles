package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ArticleId    uuid.UUID      `gorm:"primaryKey;column:article_id;type:varchar(36)"`
	ArticleTitle string         `gorm:"index;column:article_title;type:varchar(100)"`
	ArticleBody  string         `gorm:"column:article_body"`
	Category     string         `gorm:"column:category"`
	IsActive     *int           `gorm:"column:is_active;default:1"`
	Thumbnail    string         `gorm:"column:thumbnail"`
	CreatedBy    string         `gorm:"column:created_by"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedBy    string         `gorm:"column:updated_by"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedBy    string         `gorm:"column:deleted_by"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Article) TableName() string {
	return "tb_article"
}
