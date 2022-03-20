package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smallretardedfish/cakes/internal/domain"
	"log"
	"net/http"
)

type UserService interface {
	Create(inp *domain.UserSignUpInput) error
	Get(id uuid.UUID) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uuid.UUID) error
	ReadAll() ([]*domain.User, error)
}
type UserHandler struct {
	UserService UserService
}

func NewUserHandler(service UserService) *UserHandler {
	h := &UserHandler{UserService: service}
	h.initRoutes()
	return h
}

func (uh *UserHandler) Create(c *gin.Context) {
	inp := domain.UserSignUpInput{}
	err := c.BindJSON(&inp)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = uh.UserService.Create(&inp)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	c.Status(http.StatusOK)
}
func (uh *UserHandler) ReadAll(c *gin.Context) {
	users, err := uh.UserService.ReadAll()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) initRoutes() *gin.Engine {
	r := gin.Default()
	r.Group("/user")
	{
		r.POST("/sign-up", uh.Create)
		r.GET("/all", uh.ReadAll)
	}
	return r
}
