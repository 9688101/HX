package v1

import (
	"fmt"
	"net/http"

	"github.com/9688101/HX/internal/entity"
	"github.com/gin-gonic/gin"
)

func RelayNotFound(c *gin.Context) {
	err := entity.Error{
		Message: fmt.Sprintf("Invalid URL (%s %s)", c.Request.Method, c.Request.URL.Path),
		Type:    "invalid_request_error",
		Param:   "",
		Code:    "",
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": err,
	})
}
