package user

import (
	"net/http"
	"strconv"

	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user/dto"
	"github.com/ayenalhoquegit/go-gin-project-layout/pkg/response"
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
		response.RespondError(err.Code, err.Err, ctx)
	}
	response.Respond(http.StatusOK, user, ctx)

}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var newUser dto.UserDto
	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}
	//fmt.Println("New user : ", newUser)
	user, err := h.service.CreateUser(&newUser)
	if err != nil {
		response.RespondError(err.Code, err.Err, ctx)
		return
	}
	response.Respond(http.StatusCreated, user, ctx)
}
func (h *Handler) FindUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	user, err := h.service.FindUser(userId)
	if err != nil {
		response.RespondError(err.Code, err.Err, ctx)
		return
	}
	response.Respond(http.StatusOK, user, ctx)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	var newUser dto.UserDto
	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}
	user, err := h.service.UpdateUser(userId, &newUser)
	if err != nil {
		response.RespondError(err.Code, err.Err, ctx)
		return
	}
	response.Respond(http.StatusOK, user, ctx)
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	affectedRows, err := h.service.DeleteUser(userId)
	if err != nil {
		response.RespondError(err.Code, err.Err, ctx)
		return
	}
	response.Respond(http.StatusOK, affectedRows, ctx)

}
