package server

import (
	"REST-API-BookCatalog-Gin/handler"
	"REST-API-BookCatalog-Gin/repository"
	"REST-API-BookCatalog-Gin/usecase"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiServer struct {
	DB        *sql.DB
	Router    *gin.Engine
	validator *validator.Validate
}

func NewServer(db *sql.DB, validator *validator.Validate) *ApiServer {
	r := gin.New()
	return &ApiServer{
		DB:        db,
		Router:    r,
		validator: validator,
	}
}

func (server *ApiServer) ListenAndServer(port string) {
	server.registerRouter()

	http.ListenAndServe(":"+port, server.Router)
}

func (server *ApiServer) registerRouter() {
	repo := repository.NewBookRepository(server.DB)
	uCase := usecase.NewBookUsecase(repo)
	bookHandler := handler.NewBookHandler(uCase, server.validator)

	log.Println(bookHandler)
}
