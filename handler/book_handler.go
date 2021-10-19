package handler

import (
	"REST-API-BookCatalog-Gin/transport"
	"REST-API-BookCatalog-Gin/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	usecase   usecase.BookUsecase
	validator *validator.Validate
}

func NewBookHandler(bookUseCase usecase.BookUsecase, validator *validator.Validate) *bookHandler {
	return &bookHandler{
		usecase:   bookUseCase,
		validator: validator,
	}
}

func (b *bookHandler) GetList(gc *gin.Context) {
	gc.Header("Content-Type", "application/json")
	result, err := b.usecase.GetList()

	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	gc.JSON(http.StatusOK, result)
}

func (b *bookHandler) GetByID(gc *gin.Context) {
	gc.Header("Content-Type", "application/json")
	id, _ := strconv.Atoi(gc.Param("bookID"))

	result, err := b.usecase.GetByID(id)
	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	gc.JSON(http.StatusOK, result)
}

func (b *bookHandler) AddBook(gc *gin.Context) {
	gc.Header("Content-Type", "application/json")

	var requestBook transport.InsertBook
	if err := gc.ShouldBindJSON(&requestBook); err != nil {
		gc.JSON(http.StatusBadRequest, transport.ResponseError{
			Message: "error while decode request body",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// checking validation
	errorValidation := b.validator.Struct(requestBook)
	if errorValidation != nil {
		gc.JSON(http.StatusBadRequest, transport.ResponseError{
			Message: errorValidation.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	result, responseError := b.usecase.AddBook(requestBook)
	if responseError != nil {
		gc.JSON(responseError.Status, responseError)
		return
	}

	gc.JSON(http.StatusOK, result)
}

func (b *bookHandler) UpdateBook(gc *gin.Context) {
	gc.Header("Content-Type", "application/json")
	id, _ := strconv.Atoi(gc.Param("bookID"))

	var requestBook transport.UpdateBook
	if err := gc.ShouldBindJSON(&requestBook); err != nil {
		gc.JSON(http.StatusBadRequest, transport.ResponseError{
			Message: "error while decode request body",
			Status:  http.StatusBadRequest,
		})
		return
	}

	// checking validation
	errorValidation := b.validator.Struct(requestBook)
	if errorValidation != nil {
		gc.JSON(http.StatusBadRequest, transport.ResponseError{
			Message: errorValidation.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	result, responseError := b.usecase.UpdateBook(id, requestBook)
	if responseError != nil {
		gc.JSON(responseError.Status, responseError)
		return
	}

	gc.JSON(http.StatusOK, result)
}

func (b *bookHandler) DeleteBook(gc *gin.Context) {
	gc.Header("Context-Type", "application/json")
	id, _ := strconv.Atoi(gc.Param("bookID"))

	result, err := b.usecase.DeleteBook(id)
	if err != nil {
		gc.JSON(http.StatusBadRequest, err)
		return
	}

	gc.JSON(http.StatusOK, result)
}
