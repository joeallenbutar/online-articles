package impl

import (
	"context"
	"errors"
	"online-articles/entity"
	"online-articles/exception"
	"online-articles/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewArticleRepositoryImpl(DB *gorm.DB) repository.ArticleRepository {
	return &articleRepositoryImpl{DB: DB}
}

type articleRepositoryImpl struct {
	*gorm.DB
}

func (repository *articleRepositoryImpl) InsertArticle(ctx context.Context, article entity.Article) entity.Article {
	article.ArticleId = uuid.New()
	err := repository.DB.WithContext(ctx).Create(&article).Error
	exception.PanicLogging(err)
	return article
}

func (repository *articleRepositoryImpl) UpdateArticle(ctx context.Context, article entity.Article) entity.Article {
	err := repository.DB.WithContext(ctx).Where("article_id = ?", article.ArticleId).Updates(&article).Error
	exception.PanicLogging(err)
	return article
}

func (repository *articleRepositoryImpl) DeleteArticle(ctx context.Context, article entity.Article) {
	err := repository.DB.WithContext(ctx).Model(&article).Update("deleted_by", article.DeletedBy).Error
	exception.PanicLogging(err)

	err = repository.DB.WithContext(ctx).Delete(&article).Error
	exception.PanicLogging(err)
}

func (repository *articleRepositoryImpl) GetArticleById(ctx context.Context, id string) (entity.Article, error) {
	var article entity.Article
	result := repository.DB.WithContext(ctx).Unscoped().Where("article_id = ? AND deleted_at IS NULL", id).First(&article)
	if result.RowsAffected == 0 {
		return entity.Article{}, errors.New("article Not Found")
	}
	return article, nil
}

func (repository *articleRepositoryImpl) GetAllArticle(ctx context.Context) []entity.Article {
	var articles []entity.Article
	repository.DB.WithContext(ctx).Find(&articles)
	return articles
}
