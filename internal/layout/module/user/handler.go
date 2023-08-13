package user

import (
	"log"
	"net/http"

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
