package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ShowError(c *gin.Context, err error) {
	logrus.WithError(err).Error()
	c.HTML(
		http.StatusBadRequest,
		"error.html.tmpl",
		gin.H{
			"error": err.Error(),
		},
	)
	c.Abort()
}
