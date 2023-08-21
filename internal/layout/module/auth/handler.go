package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/auth/dto"
	"github.com/ayenalhoquegit/go-gin-project-layout/internal/layout/module/user"
	"github.com/ayenalhoquegit/go-gin-project-layout/pkg/config"
	"github.com/ayenalhoquegit/go-gin-project-layout/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Service *user.Service
}

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
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": u.Id,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, errr := token.SignedString([]byte(config.GetEnvValue("JWT_SECRET")))
	if errr != nil {
		fmt.Println("error", errr)
		ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Error generation token"})
		return
	}
	response.Respond(http.StatusCreated, map[string]string{"user": u.Name, "token": tokenString}, ctx)
}
