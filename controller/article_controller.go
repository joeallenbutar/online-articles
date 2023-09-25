package controller

import (
	"online-articles/common"
	"online-articles/configuration"
	"online-articles/middleware"
	"online-articles/model"
	"online-articles/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ArticleController struct {
	service.ArticleService
	configuration.Config
}

func NewArticleController(articleService *service.ArticleService, config configuration.Config) *ArticleController {
	return &ArticleController{ArticleService: *articleService, Config: config}
}

func (controller ArticleController) Route(app *fiber.App) {
	app.Post("/v1/api/article", middleware.AuthenticateJWT(controller.Config), controller.CreateArticle)
	app.Put("/v1/api/article/:id", middleware.AuthenticateJWT(controller.Config), controller.UpdateArticle)
	app.Delete("/v1/api/article/:id", middleware.AuthenticateJWT(controller.Config), controller.DeleteArticle)
	app.Get("/v1/api/article/:id", middleware.AuthenticateJWT(controller.Config), controller.GetArticleById)
	app.Get("/v1/api/article", middleware.AuthenticateJWT(controller.Config), controller.GetAllArticle)
}

// Create func CreateArticle.
// @Description CreateArticle.
// @Summary CreateArticle
// @Tags Article
// @Accept multipart/form-data
// @Produce json
// @Param imageUpload formData file true "imageUpload"
// @Param article_title formData string true "ArticleTitle"
// @Param article_body formData string true "ArticleBody"
// @Param category formData string true "Category"
// @Param is_active formData string true "isActive"
// @Param request body model.ArticleModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/article [post]
func (controller ArticleController) CreateArticle(c *fiber.Ctx) error {
	var request model.ArticleCreateModel

	if c.FormValue("article_title") == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "article_title required",
		})
	}
	if c.FormValue("article_body") == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "article_body required",
		})
	}
	if c.FormValue("category") == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "category required",
		})
	}
	if c.FormValue("is_active") == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "is_active required",
		})
	}

	imageUpload, err := c.FormFile("imageUpload")
	if imageUpload == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "imageUpload required",
		})
	}
	url, _, err := common.UploadFileToMinIO(c, imageUpload, controller.Config)
	if err != nil {
		// Return status 500 and minio connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	isActiveInt, _ := strconv.Atoi(c.FormValue("is_active"))
	isActive := &isActiveInt

	request = model.ArticleCreateModel{
		ArticleTitle: c.FormValue("article_title"),
		ArticleBody:  c.FormValue("article_body"),
		Category:     c.FormValue("category"),
		IsActive:     isActive,
		Thumbnail:    url,
		CreatedBy:    c.Locals("username").(string),
	}

	response := controller.ArticleService.CreateArticle(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// Update func UpdateArticle.
// @Description UpdateArticle.
// @Summary UpdateArticle
// @Tags Article
// @Accept multipart/form-data
// @Produce json
// @Param imageUpload formData file true "imageUpload"
// @Param article_title formData string true "ArticleTitle"
// @Param article_body formData string true "ArticleBody"
// @Param category formData string true "Category"
// @Param is_active formData string true "isActive"
// @Param request body model.ArticleModel true "Request Body"
// @Param id path string true "ArticleId"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/article/{id} [put]
func (controller ArticleController) UpdateArticle(c *fiber.Ctx) error {
	var request model.ArticleUpdateModel
	id := c.Params("id")

	if c.FormValue("article_title") == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "article_title required",
		})
	}
	if c.FormValue("article_body") == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "article_body required",
		})
	}
	if c.FormValue("category") == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "category required",
		})
	}
	if c.FormValue("is_active") == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "is_active required",
		})

	}

	imageUpload, err := c.FormFile("imageUpload")
	if imageUpload == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "imageUpload required",
		})
	}
	url, _, err := common.UploadFileToMinIO(c, imageUpload, controller.Config)
	if err != nil {
		// Return status 500 and minio connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	isActiveInt, _ := strconv.Atoi(c.FormValue("is_active"))
	isActive := &isActiveInt

	request = model.ArticleUpdateModel{
		ArticleTitle: c.FormValue("article_title"),
		ArticleBody:  c.FormValue("article_body"),
		Category:     c.FormValue("category"),
		IsActive:     isActive,
		Thumbnail:    url,
		UpdatedBy:    c.Locals("username").(string),
	}

	response := controller.ArticleService.UpdateArticle(c.Context(), request, id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// Delete func DeleteArticle.
// @Description DeleteArticle.
// @Summary DeleteArticle
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "ArticleId"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/article/{id} [delete]
func (controller ArticleController) DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	userName := c.Locals("username").(string)

	controller.ArticleService.DeleteArticle(c.Context(), id, userName)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

// FindById func GetArticleById.
// @Description GetArticleById.
// @Summary GetArticleById
// @Tags Article
// @Accept json
// @Produce json
// @Param id path string true "ArticleId"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/article/{id} [get]
func (controller ArticleController) GetArticleById(c *fiber.Ctx) error {
	id := c.Params("id")

	result := controller.ArticleService.GetArticleById(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// FindAll func GetAllArticle.
// @Description GetAllArticle.
// @Summary GetAllArticle
// @Tags Article
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/article [get]
func (controller ArticleController) GetAllArticle(c *fiber.Ctx) error {
	result := controller.ArticleService.GetAllArticle(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
