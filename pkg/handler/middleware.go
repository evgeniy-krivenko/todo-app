package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headersPart := strings.Split(header, " ")
	if len(headersPart) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "wrong format of authorization header")
		return
	}

	UserId, err := h.services.ParseToken(headersPart[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	// Записываем id пользователя в контекст
	c.Set(userCtx, UserId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("not found userId in context")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("not found userId in context")
	}
	return idInt, nil
}
