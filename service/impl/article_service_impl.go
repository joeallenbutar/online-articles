package impl

import (
	"context"
	"online-articles/common"
	"online-articles/entity"
	"online-articles/exception"
	"online-articles/model"
	"online-articles/repository"
	"online-articles/service"
	"time"

	"github.com/google/uuid"
)

func NewArticleServiceImpl(articleRepository *repository.ArticleRepository) service.ArticleService {
	return &articleServiceImpl{
		ArticleRepository: *articleRepository,
	}
}

type articleServiceImpl struct {
	repository.ArticleRepository
}

func (service *articleServiceImpl) CreateArticle(ctx context.Context, articleModel model.ArticleCreateModel) model.ArticleCreateModel {
	article := entity.Article{
		ArticleTitle: articleModel.ArticleTitle,
		ArticleBody:  articleModel.ArticleBody,
		Category:     articleModel.Category,
		IsActive:     articleModel.IsActive,
		Thumbnail:    articleModel.Thumbnail,
		CreatedBy:    articleModel.CreatedBy,
		CreatedAt:    time.Now(),
	}

	result := service.ArticleRepository.InsertArticle(ctx, article)
	common.NewLogger().Info("ArticleService-CreateArticle - ArticleId: [", result.ArticleId, "]")
	articleModel = model.ArticleCreateModel{
		ArticleId:    result.ArticleId,
		ArticleTitle: result.ArticleTitle,
		ArticleBody:  result.ArticleBody,
		Category:     result.Category,
		IsActive:     result.IsActive,
		Thumbnail:    result.Thumbnail,
		CreatedBy:    result.CreatedBy,
		CreatedAt:    result.CreatedAt,
	}
	return articleModel
}

func (service *articleServiceImpl) UpdateArticle(ctx context.Context, articleModel model.ArticleUpdateModel, id string) model.ArticleUpdateModel {
	_, err := service.ArticleRepository.GetArticleById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	article := entity.Article{
		ArticleId:    uuid.MustParse(id),
		ArticleTitle: articleModel.ArticleTitle,
		ArticleBody:  articleModel.ArticleBody,
		Category:     articleModel.Category,
		IsActive:     articleModel.IsActive,
		Thumbnail:    articleModel.Thumbnail,
		UpdatedBy:    articleModel.UpdatedBy,
		UpdatedAt:    time.Now(),
	}

	result := service.ArticleRepository.UpdateArticle(ctx, article)
	common.NewLogger().Info("ArticleService-UpdateArticle - ArticleId: [", result.ArticleId, "]")
	articleModel = model.ArticleUpdateModel{
		ArticleId:    result.ArticleId,
		ArticleTitle: result.ArticleTitle,
		ArticleBody:  result.ArticleBody,
		Category:     result.Category,
		IsActive:     result.IsActive,
		Thumbnail:    result.Thumbnail,
		UpdatedBy:    result.UpdatedBy,
		UpdatedAt:    result.UpdatedAt,
	}
	return articleModel
}

func (service *articleServiceImpl) DeleteArticle(ctx context.Context, id string, userName string) {
	article, err := service.ArticleRepository.GetArticleById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	article.DeletedBy = userName
	service.ArticleRepository.DeleteArticle(ctx, article)
	common.NewLogger().Info("ArticleService-DeleteArticle - ArticleId: [", id, "]")
}

func (service *articleServiceImpl) GetArticleById(ctx context.Context, id string) model.ArticleModel {
	getById, err := service.ArticleRepository.GetArticleById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	common.NewLogger().Info("ArticleService-GetArticleById - ArticleId: [", id, "]")

	return model.ArticleModel{
		ArticleId:    getById.ArticleId.String(),
		ArticleTitle: getById.ArticleTitle,
		ArticleBody:  getById.ArticleBody,
		Category:     getById.Category,
		IsActive:     getById.IsActive,
		Thumbnail:    getById.Thumbnail,
		CreatedBy:    getById.CreatedBy,
		CreatedAt:    getById.CreatedAt,
		UpdatedBy:    getById.UpdatedBy,
		UpdatedAt:    getById.UpdatedAt,
	}
}

func (service *articleServiceImpl) GetAllArticle(ctx context.Context) (responses []model.ArticleModel) {
	articles := service.ArticleRepository.GetAllArticle(ctx)
	for _, article := range articles {
		responses = append(responses, model.ArticleModel{
			ArticleId:    article.ArticleId.String(),
			ArticleTitle: article.ArticleTitle,
			ArticleBody:  article.ArticleBody,
			Category:     article.Category,
			IsActive:     article.IsActive,
			Thumbnail:    article.Thumbnail,
			CreatedAt:    article.CreatedAt,
		})
	}
	if len(articles) == 0 {
		return []model.ArticleModel{}
	}
	common.NewLogger().Info("ArticleService-GetAllArticle")

	return responses
}
