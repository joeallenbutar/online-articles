package repository

import (
	"context"
	"online-articles/entity"
)

type ArticleRepository interface {
	InsertArticle(ctx context.Context, article entity.Article) entity.Article
	UpdateArticle(ctx context.Context, article entity.Article) entity.Article
	DeleteArticle(ctx context.Context, article entity.Article)
	GetArticleById(ctx context.Context, id string) (entity.Article, error)
	GetAllArticle(ctx context.Context) []entity.Article
}
