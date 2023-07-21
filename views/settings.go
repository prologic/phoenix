package views

import (
	"github.com/gin-gonic/gin"
	"github.com/ordinary-dev/phoenix/backend"
	"gorm.io/gorm"
	"net/http"
)

func ShowSettings(c *gin.Context, db *gorm.DB) {
	if err := RequireAuth(c, db); err != nil {
		return
	}

	// Get a list of groups with links
	var groups []backend.Group
	result := db.
		Model(&backend.Group{}).
		Preload("Links").
		Find(&groups)

	if result.Error != nil {
		ShowError(c, result.Error)
		return
	}

	c.HTML(http.StatusOK, "settings.html.tmpl", gin.H{
		"groups": groups,
	})
}
