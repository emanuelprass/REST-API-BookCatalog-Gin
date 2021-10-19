package usecase

import (
	"REST-API-BookCatalog-Gin/entity"
	"REST-API-BookCatalog-Gin/repository"
	"REST-API-BookCatalog-Gin/transport"
	"database/sql"
	"net/http"
)

type BookUsecase interface {
	GetList() (*transport.GetList, *transport.ResponseError)
	GetByID(id int) (*transport.GetBookResponse, *transport.ResponseError)
	AddBook(data transport.InsertBook) (*transport.GeneralResponse, *transport.ResponseError)
	UpdateBook(id int, data transport.UpdateBook) (*transport.GeneralResponse, *transport.ResponseError)
	DeleteBook(id int) (*transport.GeneralResponse, *transport.ResponseError)
}

type bookUsecase struct {
	repository repository.BookRepository
}

func NewBookUsecase(bookRepository repository.BookRepository) BookUsecase {
	return &bookUsecase{
		repository: bookRepository,
	}
}

func (b *bookUsecase) GetList() (*transport.GetList, *transport.ResponseError) {
	result, err := b.repository.GetList()

	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}

	return &transport.GetList{
		Count:    len(result),
		ListBook: result,
	}, nil
}

func (b *bookUsecase) GetByID(id int) (*transport.GetBookResponse, *transport.ResponseError) {
	result, err := b.repository.GetByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			responseError := &transport.ResponseError{
				Message: "Data not found, please check your request.",
				Status:  http.StatusNotFound,
			}
			return nil, responseError
		}
	}

	return &transport.GetBookResponse{
		Data: *result,
	}, nil
}

func (b *bookUsecase) AddBook(data transport.InsertBook) (*transport.GeneralResponse, *transport.ResponseError) {
	createPayload := entity.Book{
		Name:    data.Name,
		Creator: data.Creator,
	}

	err := b.repository.AddBook(createPayload)
	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}

	result := &transport.GeneralResponse{
		Message: "Success to insert book : " + createPayload.Name + " into database",
	}

	return result, nil
}

func (b *bookUsecase) UpdateBook(id int, data transport.UpdateBook) (*transport.GeneralResponse, *transport.ResponseError) {
	result, errBook := b.repository.GetByID(id)

	if errBook != nil {
		responseError := &transport.ResponseError{
			Message: "Data not found, please check your request.",
			Status:  http.StatusNotFound,
		}
		return nil, responseError
	}

	if data.Name == "" {
		data.Name = result.Name
	}

	if data.Creator == "" {
		data.Creator = result.Creator
	}

	createPayload := entity.Book{
		Id:      id,
		Name:    data.Name,
		Creator: data.Creator,
	}

	err := b.repository.UpdateBook(createPayload)
	if err != nil {
		response := &transport.ResponseError{
			Message: "Un Processable Entity",
			Status:  http.StatusUnprocessableEntity,
		}
		return nil, response
	}

	res := &transport.GeneralResponse{
		Message: "Success to update book : " + createPayload.Name + ".",
	}

	return res, nil
}

func (b *bookUsecase) DeleteBook(id int) (*transport.GeneralResponse, *transport.ResponseError) {
	_, err := b.repository.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			responseError := &transport.ResponseError{
				Message: "Data not found, please check yout request",
				Status:  http.StatusNotFound,
			}
			return nil, responseError
		}
	}
	er := b.repository.DeleteBook(id)
	if er != nil {
		responseError := &transport.ResponseError{
			Message: "Can't delete data",
			Status:  http.StatusConflict,
		}
		return nil, responseError
	}

	response := &transport.GeneralResponse{
		Message: "Data deleted",
		// Status : http.StatusAccepted,
	}
	return response, nil
}
