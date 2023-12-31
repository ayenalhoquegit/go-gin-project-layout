package auth

import (
	"net/http"

	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/auth/dto"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user"
	"github.com/ayenalhoquegit/go-gin-project-layout/pkg/jwtpkg"
	"github.com/ayenalhoquegit/go-gin-project-layout/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Service *user.Service
	Payload jwtpkg.Payload
}

// type jwtPayload struct {
// 	Payload *jwtpkg.Payload
// }

func NewHandler(s *user.Service) *Handler {
	h := new(Handler)
	h.Service = s
	return h
}

func (h *Handler) Login(ctx *gin.Context) {
	var loginUser dto.AuthDto
	if err := ctx.BindJSON(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		return
	}
	u, err := h.Service.FindUserByEmail(loginUser.Email)
	if err != nil {
		response.RespondError(err.Code, err.Err, ctx)
		return
	}
	errr := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginUser.Password))
	if errr != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid password"})
		return
	}
	h.Payload.Id = u.Id
	tokenString := jwtpkg.GenerateToken(h.Payload)
	response.Respond(http.StatusCreated, map[string]string{"user": u.Name, "token": tokenString}, ctx)
}
