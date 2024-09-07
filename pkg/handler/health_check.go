package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Status": "OK"})
}
