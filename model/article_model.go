package model

import (
	"time"

	"github.com/google/uuid"
)

type ArticleModel struct {
	ArticleId    string    `json:"article_id"`
	ArticleTitle string    `json:"article_title"`
	ArticleBody  string    `json:"article_body"`
	Category     string    `json:"category"`
	IsActive     *int      `json:"is_active"`
	Thumbnail    string    `json:"thumbnail"`
	CreatedBy    string    `json:"created_by,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedBy    string    `json:"updated_by,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

type ArticleCreateModel struct {
	ArticleId    uuid.UUID `json:"article_id"`
	ArticleTitle string    `json:"article_title" validate:"required"`
	ArticleBody  string    `json:"article_body" validate:"required"`
	Category     string    `json:"category" validate:"required"`
	IsActive     *int      `json:"is_active" validate:"required"`
	Thumbnail    string    `json:"thumbnail" validate:"required"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
}

type ArticleUpdateModel struct {
	ArticleId    uuid.UUID `json:"article_id"`
	ArticleTitle string    `json:"article_title" validate:"required"`
	ArticleBody  string    `json:"article_body" validate:"required"`
	Category     string    `json:"category" validate:"required"`
	IsActive     *int      `json:"is_active" validate:"required"`
	Thumbnail    string    `json:"thumbnail" validate:"required"`
	UpdatedBy    string    `json:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at"`
}
