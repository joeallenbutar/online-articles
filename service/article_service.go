package service

import (
	"context"
	"online-articles/model"
)

type ArticleService interface {
	CreateArticle(ctx context.Context, model model.ArticleCreateModel) model.ArticleCreateModel
	UpdateArticle(ctx context.Context, newsModel model.ArticleUpdateModel, id string) model.ArticleUpdateModel
	DeleteArticle(ctx context.Context, id string, userName string)
	GetArticleById(ctx context.Context, id string) model.ArticleModel
	GetAllArticle(ctx context.Context) []model.ArticleModel
}
