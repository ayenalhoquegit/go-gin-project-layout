package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user/dto"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	h := new(Handler)
	h.service = s
	return h

}

func (h *Handler) FindAllUser(ctx *gin.Context) {
	user, err := h.service.findAll()
	if err != nil {
		log.Fatal(err)
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var newUser dto.UserDto
	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}
	user, err := h.service.CreateUser(&newUser)
	if err != nil {
		log.Fatal(err)
	}
	ctx.IndentedJSON(http.StatusOK, user)
}
func (h *Handler) FindUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	user, err := h.service.FindUser(userId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	affect, err := h.service.DeleteUser(userId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
	if affect > 0 {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted"})
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
}
